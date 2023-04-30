package widgets

type Executor interface {
	Encode()
	Copy()
}


type Driver interface {
	PreExecution()
}

type WidgetDriver struct {

}


type WidgetSpecification struct {
	Parameters map[string]ExecutionParameter
	Inputs map[string]ChannelParameter
	Outputs map[string]ChannelParameter
}


type ExecutorSpecification struct {

}

func (e *ExecutorSpecification) Encode() {

}

func (e *ExecutorSpecification) Copy() {

}




type ExecutionParameter struct {

}

type ChannelParameter struct {

}