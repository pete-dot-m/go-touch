package main

import (
	"fmt"
	"os"
	"pete-dot-m/go-touch/app"
)

func main() {
	files := os.Args[1:]
	if err := app.GoTouch(files); err != nil {
		fmt.Println(err.Error())
	}
}
