package dag

import (
	"fmt"

	"github.com/zoracloud/pipelines/backend/src/apiserver/workflow"
)

type Runner interface {
	Run(w workflow.LogicalWorkflow)
}

type DynamicRunner struct {
	
}

func (DynamicRunner) Run(workflow workflow.LogicalWorkflow) {
	fmt.Println("Running Workflow")
	fmt.Println("Executing workflow", workflow)
	for index , widget := range(workflow.Widgets) {
		fmt.Println("Processing Widget", widget, index)
		// fmt.Printf("Type %T \n",  widget)
		// fmt.Printf("Widget %v", widget)
		// widget.
		widget.Initialize()
		// widget.Executor()
		// a := widget.ExecutionParams()
		// fmt.Println("Execution Params", a)
		// fmt.Println(widget.Executor())
	}
}