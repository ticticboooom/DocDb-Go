package main

import (
	"DocDb-Go/wsdon"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("/home/kyle/Documents/wsdonfile.wsdon")
	fileString := string(file)
	wsdonItem := wsdon.ParseWsdon(fileString)
	print(wsdonItem)
}
