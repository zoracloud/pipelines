package widgets

import (
	"fmt"
)

type SshWidget struct {
	Name string
	Widget
	// Specification spec.WidgetSpecification
	// SshSpecification
}

type SendEmailWidget struct {
	Name string
	// SshExecutor
	// SshSpecification
}

func (s SshWidget) Executor() {
	fmt.Println("Making SSh Now ----")
	// return 10
}

func (s SendEmailWidget) Executor() {
	fmt.Println("Sending Email Now ----")
	// return 10
}

