package qmysql

import (
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/lz01wcy/qtools/qpassword"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	defaultMaxIdleConns    = 10
	defaultMaxOpenConns    = 10
	defaultConnMaxLifetime = 1 * time.Hour
	defaultConnMaxIdleTime = 1 * time.Hour
)

type Info struct {
	Host            string        `yaml:"host" json:"host"`                             //
	Port            string        `yaml:"port" json:"port"`                             //
	Account         string        `yaml:"account" json:"account"`                       //
	PasswordDecoded string        `yaml:"password_decoded" json:"password_decoded"`     //
	Password        string        `yaml:"password" json:"password"`                     //如果Password不为空 则不解密PasswordDecoded
	DatabaseName    string        `yaml:"database_name" json:"database_name"`           //库名称
	UsePool         bool          `yaml:"use_pool" json:"use_pool"`                     //是否使用连接池
	MaxIdleConns    int           `yaml:"max_idle_conns" json:"max_idle_conns"`         //是否最大保持连接数
	MaxOpenConns    int           `yaml:"max_open_conns" json:"max_open_conns"`         //最大打开连接数
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" json:"conn_max_lifetime"`   //连接最大生命周期
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" json:"conn_max_idle_time"` //连接最大保持时间
}

func (i *Info) DecodePassword() (err error) {
	if i.Password == "" {
		i.Password, err = qpassword.Decode(i.PasswordDecoded)
	}
	return
}

func (i *Info) Check() error {
	if i.UsePool {
		if i.MaxIdleConns <= 0 {
			i.MaxIdleConns = defaultMaxIdleConns
		}
		if i.MaxOpenConns <= 0 {
			i.MaxOpenConns = defaultMaxOpenConns
		}
		if i.ConnMaxLifetime <= 0 {
			i.ConnMaxLifetime = defaultConnMaxLifetime
		}
		if i.ConnMaxIdleTime <= 0 {
			i.ConnMaxIdleTime = defaultConnMaxIdleTime
		}
	}
	return nil
}

func NewClient(info *Info, models []any, options ...gorm.Option) (*gorm.DB, error) {
	if err := info.DecodePassword(); err != nil {
		return nil, err
	}
	if err := info.Check(); err != nil {
		return nil, err
	}

	c := mysqlDriver.NewConfig()
	c.User = info.Account
	c.Passwd = info.Password
	c.Net = "tcp"
	c.Addr = fmt.Sprintf("%s:%s", info.Host, info.Port)
	c.DBName = info.DatabaseName
	c.ParseTime = true
	c.Loc = time.Local
	// 不鼓励使用该charset参数，因为它会向服务器发出额外的查询。除非您需要回退行为，否则请改用collation。
	//charset=utf8mb4
	//Collation默认是utf8mb4_general_ci 无需设置charset=utf8mb4 等效的
	//c.Collation

	db, err := gorm.Open(mysql.Open(c.FormatDSN()), options...)
	if err != nil {
		return nil, err
	}

	// 设置连接池 是需要的吗
	if info.UsePool {
		if sqlDB, err := db.DB(); err != nil {
			return nil, err
		} else {
			//mybatis: 默认最大连接数10，最小连接数5
			//  protected int poolMaximumActiveConnections = 10;
			//  protected int poolMaximumIdleConnections = 5;
			//  protected int poolMaximumCheckoutTime = 20000;
			//  protected int poolTimeToWait = 20000;
			//  protected int poolMaximumLocalBadConnectionTolerance = 3;
			sqlDB.SetMaxIdleConns(info.MaxIdleConns)
			sqlDB.SetMaxOpenConns(info.MaxOpenConns)
			sqlDB.SetConnMaxLifetime(info.ConnMaxLifetime)
			sqlDB.SetConnMaxIdleTime(info.ConnMaxIdleTime)
		}
	}

	if len(models) > 0 {
		if err = db.Migrator().AutoMigrate(models...); err != nil {
			return nil, fmt.Errorf("建表失败!错误: %s", err.Error())
		}
	}

	return db, nil
}
