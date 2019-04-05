package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	bitbucket "github.com/nachocano/bitbucket-source/pkg/adapter"
	"gopkg.in/go-playground/webhooks.v3"
	bb "gopkg.in/go-playground/webhooks.v3/bitbucket"
)

const (
	// Environment variable containing the HTTP port
	envPort = "PORT"

	// Environment variable containing BitBucket UUID.
	envUUID = "BITBUCKET_UUID"
)

func main() {
	sink := flag.String("sink", "", "uri to send events to")

	flag.Parse()

	if sink == nil || *sink == "" {
		log.Fatalf("No sink given")
	}

	port := os.Getenv(envPort)
	if port == "" {
		port = "8080"
	}

	uuid := os.Getenv(envUUID)
	if uuid == "" {
		// TODO validate WebHook UUID is given, otherwise fail.
		// This should be done for security issues, as the library will validate the incoming events
		// correspond to this particular WebHook, or discard them otherwise.
		log.Printf("No UUID given")
	}

	log.Printf("Sink is: %q", *sink)

	ra := &bitbucket.Adapter{
		Sink: *sink,
	}

	hook := bb.New(&bb.Config{UUID: uuid})
	hook.RegisterEvents(ra.HandleEvent,
		bb.RepoPushEvent,
		bb.RepoForkEvent,
		bb.RepoUpdatedEvent,
		bb.RepoCommitCommentCreatedEvent,
		bb.RepoCommitStatusCreatedEvent,
		bb.RepoCommitStatusUpdatedEvent,
		bb.IssueCreatedEvent,
		bb.IssueUpdatedEvent,
		bb.IssueCommentCreatedEvent,
		bb.PullRequestCreatedEvent,
		bb.PullRequestUpdatedEvent,
		bb.PullRequestApprovedEvent,
		bb.PullRequestUnapprovedEvent,
		bb.PullRequestMergedEvent,
		bb.PullRequestDeclinedEvent,
		bb.PullRequestCommentCreatedEvent,
		bb.PullRequestCommentUpdatedEvent,
		bb.PullRequestCommentDeletedEvent)

	addr := fmt.Sprintf(":%s", port)
	err := webhooks.Run(hook, addr, "/")
	if err != nil {
		log.Fatalf("Failed to run the BitBucket WebHook")
	}
}
