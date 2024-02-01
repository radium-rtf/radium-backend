package app

import (
	"github.com/radium-rtf/radium-backend/config"
	radium "github.com/radium-rtf/radium-backend/internal/radium/app"
	wave "github.com/radium-rtf/radium-backend/internal/wave/app"
	"log"
	"sync"
)

type Runner interface {
	Run() error
}

func Run(cfg *config.Config) {
	db, err := openDB(cfg.PG)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	radiumApp := radium.NewApp(cfg, db)
	waveApp := wave.NewApp(cfg, db)

	runners := []Runner{radiumApp, waveApp}
	for _, runner := range runners {
		runner := runner
		go func() {
			defer wg.Done()
			err := runner.Run()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	wg.Wait()
}
