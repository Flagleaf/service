package main

import (
	"fmt"
	"service/starter"
)

func main() {
	app, err := starter.GetApplicationSingleton()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	app.Run()
}
