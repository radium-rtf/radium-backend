package main

import (
	"errors"
	"github.com/radium-rtf/radium-backend/config"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	host       string
	healthPath string
	basePath   string

	userEmail    string
	userPassword string

	attempts = 20
)

//goland:noinspection HttpUrlsUsage
func TestMain(m *testing.M) {
	cfg := config.MustConfig().HTTP

	userEmail = "test.test.test@urfu.me"
	userPassword = "passworD!123"

	host = "server:" + cfg.Port
	healthPath = "http://" + host + "/healthz"
	basePath = "http://" + host

	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var (
		err  error
		resp *http.Response
	)

	for attempts > 0 {
		resp, err = http.Get(healthPath)
		if err == nil {
			break
		}
		if attempts%3 == 0 {
			log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)
		}

		time.Sleep(time.Second * 2)
		attempts--
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("")
}
