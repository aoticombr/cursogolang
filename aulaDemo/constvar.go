package main

import (
	"github.com/aoticombr/golang/config"
)

var NomeApp = "meuapp"
var Config *config.Config

func init() {
	Config = config.NewConfig()
}
