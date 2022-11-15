package nutil

import (
	"github.com/gurkankaymak/hocon"
)

type Config struct {
	orgConf *hocon.Config
}

const (
	loadTypePath = iota
	loadTypeString
)

type LoadType int

func createConfig(str string, loadType LoadType) (*Config, error) {

	conf := &Config{}
	var realConf *hocon.Config
	var err error

	switch loadType {
	case loadTypePath:
		realConf, err = hocon.ParseResource(str)
		if err != nil {
			return nil, err
		}
		break
	case loadTypeString:
		realConf, err = hocon.ParseString(str)
		if err != nil {
			return nil, err
		}
		break
	default:
		return nil, NewBasicError("Not Supported LoadType = %d", loadType)
	}

	conf.orgConf = realConf

	return conf, nil
}

func CreateNewConfigFromString(str string) (*Config, error) {
	return createConfig(str, loadTypeString)
}

func CreateNewConfigFromPath(path string) (*Config, error) {
	return createConfig(path, loadTypePath)
}

func (r *Config) GetInt(name string) int {
	defer RecoverPanic()
	ret := r.orgConf.GetInt(name)
	return ret
}

func (r *Config) GetIntArray(name string) []int {
	defer RecoverPanic()
	ret := r.orgConf.GetIntSlice(name)
	return ret
}

func (r *Config) GetString(name string) string {
	defer RecoverPanic()
	ret := r.orgConf.GetString(name)
	return ret
}

func (r *Config) GetStringArray(name string) []string {
	defer RecoverPanic()
	ret := r.orgConf.GetStringSlice(name)
	return ret
}

func (r *Config) GetConfigObject(name string) *Config {
	defer RecoverPanic()

	ret := &Config{}
	convConf := r.orgConf.GetConfig(name)
	ret.orgConf = convConf
	return ret
}

func (r *Config) GetBool(name string) bool {
	defer RecoverPanic()
	ret := r.orgConf.GetBoolean(name)
	return ret
}
