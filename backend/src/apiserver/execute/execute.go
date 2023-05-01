package execute

import "github.com/zoracloud/pipelines/backend/src/apiserver/widgets"


type Executor interface {
	Initialize()
	ExecutionParams() widgets.WidgetSpecification
	Executor()
}