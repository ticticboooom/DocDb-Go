package main

import (
	"DocDb-Go/wsdon"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile("/home/kyle/Documents/wsdonfile.wsdon")
	fileString := string(file)
	wsdonItem := wsdon.ParseWsdon(fileString)
	stringsTings := wsdon.Stringify(wsdonItem)
	ioutil.WriteFile("/home/kyle/Documents/wsdonfile2.wsdon", []byte(stringsTings), os.ModePerm)
}
