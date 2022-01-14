package config

import "github.com/spf13/viper"

type Config struct {
  Nicloudb Nicloudb
  Page Page
}

type Page struct {
 Offset int
}

type Nicloudb struct {
  Dbname string
  Host string
  Port string
  User string
  Passwd string
}

var config Config
func Exportconfig() (*Config, error) {
  viper.SetConfigName("setting")
  viper.AddConfigPath("./conf")

  err := viper.ReadInConfig()
  if err != nil {
    return nil, err
  }
  viper.Unmarshal(&config)
  return &config, err
}
