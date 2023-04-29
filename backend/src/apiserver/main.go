package main

import (
	"github.com/zoracloud/pipelines/backend/src/apiserver/dag"
	"github.com/zoracloud/pipelines/backend/src/apiserver/execute"
	"github.com/zoracloud/pipelines/backend/src/apiserver/widgets"
	"github.com/zoracloud/pipelines/backend/src/apiserver/workflow"
)

func main() {
	runner := dag.DynamicRunner{}

	widgetA := widgets.SshWidget{Name: "SSH Widget A"}
	widgetB := widgets.SendEmailWidget{Name: "Email Widget B"}

	workflow := workflow.LogicalWorkflow{Widgets: []execute.Executor{&widgetA, &widgetB}}
	
	runner.Run(workflow)
}
