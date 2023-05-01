package widgets

import "fmt"

// type WidgetExecutor interface {
// 	Initialize(spec WidgetSpecification)
// 	Execute()
// 	RunExecutor() WidgetSpecification
// }


type  Widget struct {
	WidgetSpec WidgetSpecification
	// ExecutorSpecification
	// WidgetDriver
}



func (w Widget) Initialize(spec WidgetSpecification) {
	fmt.Println(spec)
}

func (w Widget) Execute() {

}