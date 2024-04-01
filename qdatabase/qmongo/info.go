package qmongo

import "github.com/lzyorozuya/qtools/qpassword"

type Info struct {
	Url             string `yaml:"url" json:"url"` //Url优先于其他配置
	Host            string `yaml:"host" json:"host"`
	Port            string `yaml:"port" json:"port"`
	Account         string `yaml:"account" json:"account"`
	PasswordDecoded string `yaml:"password_decoded" json:"password_decoded"`
	Password        string `yaml:"password" json:"password"`
}

func (i *Info) DecodePassword() (err error) {
	if i.Password == "" {
		i.Password, err = qpassword.Decode(i.PasswordDecoded)
	}
	return
}

func (i *Info) Check() error {

	return nil
}
