package main

import (
	"context"
	"flag"
	"log"

	"github.com/kardianos/service"
)

type (
	program struct {
		exit chan struct{}
	}
)

var (
	ctx     = context.Background()
	svcFlag = flag.String("service", "", "Control the system service (start, stop, install, uninstall)")
	logger  service.Logger
)

func (p *program) Start(s service.Service) error {
	log.Println("Starting service.")
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	log.Println("Stopping service.")
	close(p.exit)
	return nil
}

func main() {
	flag.Parse()
	appChDir()

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	svcConfig := &service.Config{
		Name:        "Mundi",
		DisplayName: "Mundi document conversion server",
		Description: "Document conversion service",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target",
		},
		Option: options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	if err != nil {
		logger.Error(err)
	}

}
