package main

import (
	"fmt"

	"github.com/zoracloud/pipelines/backend/src/apiserver/dag"
	"github.com/zoracloud/pipelines/backend/src/apiserver/widgets"
	"github.com/zoracloud/pipelines/backend/src/apiserver/workflow"
)

func main() {
	runner := dag.DynamicRunner{}

	widgetA := widgets.Widgets{Name: "SSh"}
	widgetB := widgets.Widgets{Name: "Email"}

	workflow := workflow.LogicalWorkflow{Widgets: []widgets.Widgets{widgetA, widgetB}}
	
	runner.Run(workflow)
	fmt.Println(runner, workflow)
}
