package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	ex1()
}

func ex1() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current dir: ", dir)

	d := http.Dir(".")
	f, err := d.Open("stuff/things.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
}
