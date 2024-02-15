package main

import (
	"fmt"
	"log"

	"github.com/GeorgeEngland/perry"
	"github.com/GeorgeEngland/perry/store"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	// Create the client object just once per process
	c, err := client.Dial(client.Options{HostPort: "temporal:7233"})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// UPDATE CHECKER & SENDER TO YOUR PREFERRED IMPLEMENTATION
	// UPDATE newEngine in workflow.go if a different engine implementation is desired
	// (workflows are deterministic (should not depend on external state - cannot pass in dependencies))
	activities := &perry.Activities{
		Checker: perry.NewInMemoryAlertChecker(),
		Sender:  perry.NewLoggingNotificationSender(),
		Store:   store.NewInMemoryStore(),
	}

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, perry.CheckAlertsQueue, worker.Options{})
	w.RegisterWorkflow(perry.HandleAlertsWorkflow)
	w.RegisterActivity(activities)
	// Start listening to the Task Queue
	fmt.Println("Listening to task Queue")
	err = w.Run(worker.InterruptCh())
	fmt.Println("Exiting...")
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
