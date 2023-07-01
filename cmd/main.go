package main

import (
	goose "github.com/bypepe77/goose/pkg"
)

func main() {
	mongoURI := "mongodb://localhost:27017"
	g, err := goose.NewGoose(mongoURI, "test")
	if err != nil {
		panic(err)
	}

	a := g.GetCollection("test")
}
