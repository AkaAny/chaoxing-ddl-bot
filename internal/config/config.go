package config

import "github.com/pelletier/go-toml"

type Config struct {
	Port int `toml:"port"`
	CAS  struct {
		UserName string `toml:"userName"`
		Password string `toml:"password"`
	} `toml:"CAS"`
	DB struct {
		DSN string `toml:"dsn"`
	} `toml:"DB"`
}

func Unmarshall(path string) *Config {
	tree, err := toml.LoadFile(path)
	if err != nil {
		panic(err)
	}
	serviceTree := tree.Get("Service").(*toml.Tree)
	var conf = new(Config)
	err = serviceTree.Unmarshal(conf)
	if err != nil {
		panic(err)
	}
	return conf
}
