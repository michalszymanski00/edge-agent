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
	currentVersion = "0.1.0"
)

func heartbeat() {
	hostname, _ := os.Hostname()
	log.Println("heartbeat from", hostname)
}

func maybeUpdate() {
	v := semver.MustParse(currentVersion)
	latest, err := selfupdate.UpdateSelf(v, repo)
	if err != nil {
		log.Println("self-update error:", err)
		return
	}
	if latest.Version.Equals(v) {
		log.Println("No update available")
	} else {
		log.Printf("Successfully updated to version %s", latest.Version)
	}
}

func main() {
	for range time.Tick(tick) {
		heartbeat()
		maybeUpdate()
	}
}
