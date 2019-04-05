package main

import (
	"log"

	"github.com/nachocano/bitbucket-source/pkg/apis"
	controller "github.com/nachocano/bitbucket-source/pkg/reconciler"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

func main() {
	// Get a config to talk to the apiserver.
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Registering Components.")

	// Setup Scheme for all resources
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		log.Fatal(err)
	}

	// Setup BitBucket Controller.
	if err := controller.Add(mgr); err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting BitBucket Controller.")

	// Start the Cmd
	log.Fatal(mgr.Start(signals.SetupSignalHandler()))
}
