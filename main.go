package main

import "github.com/dtylman/gowd"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	body, err := gowd.ParseElement("<h1>Hello World!</h1>", nil)
	check(err)
	gowd.Run(body)
}
