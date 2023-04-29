package widgets

import (
	"fmt"

	// "github.com/zoracloud/pipelines/backend/src/apiserver/execute"
)

type SshExecutor struct {

}

type SshSpecification struct {
	

}


type SshWidget struct {
	Name string
	// SshExecutor
	// SshSpecification
}

type SendEmailWidget struct {
	Name string
	// SshExecutor
	// SshSpecification
}

func (s *SshWidget) Executor() {
	fmt.Println("Making SSh Now ----")
	// return 10
}

func (s *SendEmailWidget) Executor() {
	fmt.Println("Sending Email Now ----")
	// return 10
}

