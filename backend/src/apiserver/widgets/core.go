package widgets

import (
	"fmt"

)

type SshSpec struct {
	WidgetSpecification
}

func (s SshSpec) ParseArguments() {

}

type SshWidget struct {
	Name string
	// WidgetSpec WidgetSpecification
	// ExecutorSpec ExecutorSpecification
	Widget
	// Specification spec.WidgetSpecification
	// SshSpecification
}


func (s SshWidget) Initialize(){
	fmt.Println("Initializing SSh --")
	parameters := make(map[string]ExecutionParameter)
	parameters["networkElements"] = ExecutionParameter{}
	s.WidgetSpec = WidgetSpecification{Parameters: parameters}
	fmt.Println("spec", s.WidgetSpec.Parameters)
	// return s.WidgetSpec
}


func (s SshWidget) Executor(){
	fmt.Println("Making SSh Now ----")
	// return 10
}

func (s SshWidget) ExecutionParams() WidgetSpecification{
	fmt.Println("Execution Params", s.WidgetSpec.Parameters)
	fmt.Println(s.WidgetSpec.Parameters)
	return s.WidgetSpec
}


type SendEmailWidget struct {
	Name string
	Widget
	// SshExecutor
	// SshSpecification
}

func (s SendEmailWidget) Initialize(){
	fmt.Println("Initializing Sending Email ---")
	// specification := SshSpec{}.ParseArguments()
	fmt.Println()
}

func (s SendEmailWidget) Executor() {
	fmt.Println("Sending Email Now ----")
	// return 10
}

