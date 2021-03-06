package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

const (
	adapterID       = "hci0"
	clientAdapterID = "hci1"

	objectName = "org.bluez"
	objectPath = "/org/bluez/example/service"

	appName = "Example golang"
)

func main() {

	log.SetLevel(log.DebugLevel)

	if len(os.Args) > 0 && os.Args[len(os.Args)-1] == "client" {
		createClient(objectName, objectPath)
	} else {
		err := registerApplication()
		if err != nil {
			panic(err)
		}
	}

	select {}
}
