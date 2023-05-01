package workflow

import (
	"github.com/zoracloud/pipelines/backend/src/apiserver/execute"
)


type LogicalWorkflow struct {
	Widgets []execute.Executor
}