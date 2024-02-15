package main

import (
	"context"
	"log"
	"time"

	"github.com/GeorgeEngland/perry"
	"github.com/GeorgeEngland/perry/data"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{HostPort: "temporal:7233"})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	ctx := context.Background()
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: perry.CheckAlertsQueue,
	}
	for {
		w, err := c.ExecuteWorkflow(context.Background(), workflowOptions, perry.HandleAlertsWorkflow,
			perry.HandleAlertsWorkflowInput{CheckTime: time.Second * 10})
		if err != nil {
			log.Fatalln("error in workflow", w)
		}
		var res []data.NotificationData
		w.Get(ctx, &res)
		time.Sleep(time.Second * 10)
	}

}
