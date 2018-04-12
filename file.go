package main

import (
  "fmt"
  "io/ioutil"
)

func main() {

  file, err := ioutil.ReadFile("main.go")

  if err != nil {
	fmt.Println(err)
	return
  }

  fmt.Println("%q", file)
}

