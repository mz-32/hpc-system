package lib

import (

  "github.com/BurntSushi/toml"
)

type Config struct {
    Server ServerConfig
}

type ServerConfig struct {
  Host  string        `toml:"host"`
  WebPort  string        `toml:"webport"`
  SocketPort string `toml:"socketport"`
  UpdateInterval int `toml:"update_interval"`
  Nodes []NodeConfig `toml:"node"`
}

type NodeConfig struct {
  Id    int     `toml:"id"`
  Host  string  `toml:"host"`
  Ip    string  `toml:"ip"`
  Port  string        `toml:"port"`
}

func GetServerConfig() (Config, []error) {
   var config Config
   var errors []error
   _, err := toml.DecodeFile("config.tml", &config)
    if err != nil {
      errors = append(errors,err)
    }
    return config, errors
}
