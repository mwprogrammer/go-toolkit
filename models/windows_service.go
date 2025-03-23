package models

import (
	"log"
	"time"

	"golang.org/x/sys/windows/svc"
)

type WindowsService struct{
	Name string
	Interval int64
	IsInDebugMode bool
	Function func()
}

func (service *WindowsService) Execute(arguments []string, r <-chan svc.ChangeRequest, status chan<- svc.Status) (bool, uint32) {

	const valid_commands = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue

	tick := time.Tick(time.Duration(service.Interval) * time.Second)

	status <- svc.Status{State: svc.StartPending}
	status <- svc.Status{State: svc.Running, Accepts: valid_commands}

loop:
	for {

		select {

		case <-tick:
			log.Print("Tick handled...")

		case c := <-r:

			switch c.Cmd {
			case svc.Interrogate:
				status <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				log.Print("Shutting service ...")
				break loop
			case svc.Pause:
				status <- svc.Status{State: svc.Paused, Accepts: valid_commands}
			case svc.Continue:
				status <- svc.Status{State: svc.Paused, Accepts: valid_commands}
			default:
				log.Printf("Unexpected service control request #%d", c)
			}
		}
	}

	status <- svc.Status{State: svc.StopPending}
	return false, 1
}