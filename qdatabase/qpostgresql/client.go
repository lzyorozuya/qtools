package qpostgresql

import (
	"fmt"
	"github.com/lz01wcy/qtools/qpassword"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Info struct {
	Host            string `yaml:"host" json:"host"`                         //
	Port            string `yaml:"port" json:"port"`                         //
	Account         string `yaml:"account" json:"account"`                   //
	PasswordDecoded string `yaml:"password_decoded" json:"password_decoded"` //
	Password        string `yaml:"password" json:"password"`                 //如果Password不为空 则不解密PasswordDecoded
	DatabaseName    string `yaml:"database_name" json:"database_name"`       //
	SSLMode         bool   `yaml:"ssl_mode" json:"ssl_mode"`                 //
	sslModeString   string ``                                                //
	TimeZone        string `yaml:"time_zone" json:"time_zone"`               //如果为空 则使用Asia/Shanghai
}

func (i *Info) DecodePassword() (err error) {
	if i.Password == "" {
		i.Password, err = qpassword.Decode(i.PasswordDecoded)
	}
	if i.SSLMode {
		i.sslModeString = "enable"
	} else {
		i.sslModeString = "disable"
	}
	if i.TimeZone == "" {
		i.TimeZone = "Asia/Shanghai"
	}
	return
}

func (i *Info) Check() error {

	return nil
}

func NewClient(info *Info, models []any, options ...gorm.Option) (*gorm.DB, error) {
	if err := info.DecodePassword(); err != nil {
		return nil, err
	}
	if err := info.Check(); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", info.Host, info.Account, info.Password, info.DatabaseName, info.Port)
	db, err := gorm.Open(postgres.Open(dsn), options...)

	if len(models) > 0 {
		if err = db.Migrator().AutoMigrate(models...); err != nil {
			return nil, fmt.Errorf("建表失败!错误: %s", err.Error())
		}
	}

	return db, nil
}
