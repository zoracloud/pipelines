package widgets

type Executor interface {
	Encode()
	Copy()
}

// type WidgetSpecification interface {
// 	ParseArguments()
// }


type Driver interface {
	PreExecution()
}

type WidgetDriver struct {

}


type WidgetSpecification struct {
	Parameters map[string]ExecutionParameter
	// Inputs map[string]ChannelParameter
	// Outputs map[string]ChannelParameter
}


// 
type ExecutorSpecification struct {

}

func (e *ExecutorSpecification) Encode() {

}

func (e *ExecutorSpecification) Copy() {

}



// ExecutionParameter An execution parameter in a WidgetSpec
// This type of parameter should be specified in the Parameter map
// of a WidgetSpec
type ExecutionParameter struct {
	
}

func (e ExecutionParameter) Initialize() {

}

type ChannelParameter struct {

}