package skelego

import (
	"github.com/spf13/viper"
)

//Config Wrapper around Viper that allows a little more customizability in declaring configs
type Config interface {
	Add(string, interface{})
	DefaultSetting(string, interface{})
	Get(string) interface{}
}

//Configuration config for any services included skelego
type configuration struct {
	services *viper.Viper
}

//NewConfig creates new config
func NewConfig(configFile string, logger Logging, dirs ...string) Config {
	v := viper.New()
	v.SetConfigName(configFile)

	if len(dirs) == 0 {
		v.AddConfigPath(".")
		v.AddConfigPath("..")
	} else {
		for _, dir := range dirs {
			v.AddConfigPath(dir)
		}
	}

	if err := v.ReadInConfig(); err != nil {
		logger.LogFatal(err.Error())
	}
	logger.LogEvent("Config file read.")
	return &configuration{
		services: v,
	}
}

//Add adds key and value to configuration store
func (c *configuration) Add(key string, value interface{}) {
	c.services.Set(key, value)
}

//DefaultSetting Default settings for any service
func (c *configuration) DefaultSetting(key string, value interface{}) {
	c.services.SetDefault(key, value)
}

//Get gets configured value from configuration
func (c *configuration) Get(key string) interface{} {
	return c.services.Get(key)
}
