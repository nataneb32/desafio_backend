package main

import (
	"nataneb32.live/hospedagem/pkg/app"
)

func main() {
	a := app.Start()
	r := a.CreateAppHandlersGin()
	r.Run()
}
