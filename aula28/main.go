package main

import (
	_ "aula28/api"

	Fmk "github.com/aoticombr/golang/framework"
)

func main() {
	Fmk.NewAppDev("aula28").Execute(Fmk.AddApi())
}
