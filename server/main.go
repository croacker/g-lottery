package main

import (
	"fmt"
	"log"

	"./conf"
	"./persistsql"
)

func main() {
	persistsql.Init()
	appConf := conf.Get()
	fmt.Println("IncomingCheckFolder", appConf.IncomingCheckFolder)

	var c = make(chan string)

	go httpserver.StartGin()

	defer filewatcher.Watch(appConf.IncomingCheckFolder, c).Close()
	doWait()
}

//Давай, жди
func doWait() {
	done := make(chan bool)
	<-done
}

//Обработать ошибку
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
