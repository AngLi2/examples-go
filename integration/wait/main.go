package main

import (
	"github.com/zenaton/examples-go/integration/client" // initialize zenaton client with credentials
	"github.com/zenaton/examples-go/workflows"
)
func init(){client.SetEnv("wait.env")}
func main() {
	workflows.WaitWorkflow.NewInstance().Dispatch()
}