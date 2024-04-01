package qredis

import (
	"fmt"
	"github.com/lz01wcy/qtools/qpassword"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host            string
	Port            string
	UseAuth         bool
	PasswordDecoded string
	password        string
	Username        string
	DatabaseIndex   int
}

func (c *Config) DecodePassword() (err error) {
	if c.UseAuth {
		c.password, err = qpassword.Decode(c.PasswordDecoded)
	}
	return
}

func NewClient(config *Config) (*redis.Client, error) {
	if err := config.DecodePassword(); err != nil {
		return nil, err
	}

	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Username: config.Username,
		Password: config.password,
		DB:       config.DatabaseIndex,
	}

	c := redis.NewClient(&opt)
	return c, nil
}
