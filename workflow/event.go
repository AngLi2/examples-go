package workflow

import (
	"github.com/zenaton/examples-go/task"
	"github.com/zenaton/zenaton-go/v1/zenaton"
)

var EventWorkflow = zenaton.NewWorkflow(zenaton.WorkflowParams{
	Name: "EventWorkflow",
	HandleFunc: func() {
		task.A.Execute()
		task.B.Execute()
	},
	OnEvent: func(eventName string, eventData interface{}) {
		if eventName == "MyEvent" {
			task.C.Execute()
		}
	},
	ID: func() string {
		return "MyId"
	},
})
