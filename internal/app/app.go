package app

import (
	"github.com/radium-rtf/radium-backend/config"
	radium "github.com/radium-rtf/radium-backend/internal/radium/app"
	wave "github.com/radium-rtf/radium-backend/internal/wave/app"
	"github.com/radium-rtf/radium-backend/pkg/closer"
	"log"
)

type app interface {
	Run() error
	Shutdown() error
}

func Run(cfg *config.Config) {
	db, err := openDB(cfg.PG)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	radiumApp := radium.NewApp(cfg, db)
	waveApp := wave.NewApp(cfg, db)

	runners := []app{radiumApp, waveApp}

	var closer = closer.New()
	var notify = make(chan error, len(runners))
	for _, runner := range runners {
		runner := runner
		go func() {
			err := runner.Run()
			if err != nil {
				notify <- err
			}
		}()
		closer.Add(runner.Shutdown)
	}

	err = <-notify
	log.Println(err)
	for len(notify) != 0 {
		err = <-notify
		log.Println(err)
	}
	closer.Close()
}
