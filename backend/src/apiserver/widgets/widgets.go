package widgets

type WidgetExecutor interface {
	Execute()
}

type  Widget struct {
	WidgetSpecification
	ExecutorSpecification
	WidgetDriver
}

func (w Widget) Execute() {

}