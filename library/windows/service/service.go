package service

import (
	"log"

	"github.com/mwprogrammer/go-utilities/models"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
)

var service models.WindowsService

func SetName(name string) {
	service.Name = name
}

func SetFunction(function func()) {
	service.Function = function
}

func SetInterval(seconds int64) {
	service.Interval = seconds
}

func UseDebugMode() {
	service.IsInDebugMode = true
}

func Run() {

	service.Function()
	
	if service.IsInDebugMode {

		err := debug.Run(service.Name, &service)

		if err != nil {
			log.Fatalln("Error running service in Debug Mode.")
		}

	} else {

		err := svc.Run(service.Name, &service)

		if err != nil {
			log.Fatalln("Error running service in Service Control Mode.")
		}

	}

}