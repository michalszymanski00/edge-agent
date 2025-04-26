package main

import (
	"log"
	"os"
	"time"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const (
	repo           = "michalszymanski00/edge-agent"
	tick           = 30 * time.Second
	currentVersion = "0.1.3"
)

func heartbeat() {
	hostname, _ := os.Hostname()
	log.Println("heartbeat from", hostname)
}

func maybeUpdate() {
	v := semver.MustParse(currentVersion)

	var cfg selfupdate.Config
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		cfg.APIToken = token
	}

	updater, err := selfupdate.NewUpdater(cfg)
	if err != nil {
		log.Println("failed to create updater:", err)
		return
	}

	latest, err := updater.UpdateSelf(v, repo)
	if err != nil {
		log.Println("self-update error:", err)
		return
	}

	if latest.Version.Equals(v) {
		log.Println("No update available")
	} else {
		log.Printf("Successfully updated to version %s", latest.Version)
		os.Exit(0) // Kubernetes will restart the pod automatically
	}
}

func main() {
	for range time.Tick(tick) {
		heartbeat()
		maybeUpdate()
	}
}
