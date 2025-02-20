//go:generate openssl genrsa -out ./certs/server.key 2048
//go:generate openssl ecparam -genkey -name secp384r1 -out ./certs/server.key
//go:generate openssl req -new -x509 -sha256 -key ./certs/server.key -out ./certs/server.crt -days 3650 -subj "/C=RU"
package main

import (
	"fmt"
	"github.com/wissance/stringFormatter"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	osSignal := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	app := CreateAppWithConfigs("./config.json", "./data.json", "./keyfile")
	res, initErr := app.Init()
	if initErr != nil {
		fmt.Println("An error occurred during app init, terminating the app")
		os.Exit(-1)
	} else {
		fmt.Println("Application was successfully initialized")
	}

	res, err := app.Start()
	if !res {
		msg := stringFormatter.Format("An error occurred during starting application, error is: {0}", err.Error())
		fmt.Println(msg)
	} else {
		fmt.Println("Application was successfully started")
	}

	go func() {
		sig := <-osSignal
		//logging.InfoLog(stringFormatter.Format("Got signal from OS: {0}", sig))
		fmt.Println(stringFormatter.Format("Got signal from OS: \"{0}\", stopping", sig))
		done <- true
	}()
	<-done

	res, err = app.Stop()
	if !res {
		msg := stringFormatter.Format("An error occurred during stopping application, error is: {0}", err.Error())
		fmt.Println(msg)
	} else {
		fmt.Println("Application was successfully stopped")
	}

}
