package boostrap

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

var conf = make(map[string]interface{})

func LoadConf() error {
	pwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "get current directory fail")
	}
	confPath := path.Join(pwd, "conf/db.yaml")
	yfile, err := ioutil.ReadFile(confPath)
	if err != nil {
		return errors.Wrap(err, "read conf fil fail")
	}
	err2 := yaml.Unmarshal(yfile, &conf)
	if err2 != nil {
		return errors.Wrap(err, "unmarshal yaml file fail")
	}
	return nil
}

func GetConf() *map[string]interface{}{
	return &conf
}