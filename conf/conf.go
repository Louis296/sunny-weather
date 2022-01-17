package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func GetConf() (Conf, error) {
	file, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		return Conf{}, err
	}
	var ans Conf
	err = yaml.UnmarshalStrict(file, &ans)
	if err != nil {
		return Conf{}, err
	}
	return ans, nil
}
