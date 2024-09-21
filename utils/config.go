package utils

import "github.com/spf13/viper"

// Config stores all the configurations of the app and is user by viper
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string	`mapstructure:"DB_SOURCE"` 
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`	
}

// LoadConfig readsconfiguration from files
func LoadConfig(path string)(config Config,err error){
	viper.AddConfigPath(path)
	// viper.SetConfigName("../app.env")
	viper.SetConfigFile(path +"app.env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}