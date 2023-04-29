package launcher

// Responsible for launching driver executor and publisher of a component
type Launcher interface {
	// Execute the component, includes driver, executor and publisher. 
	// Returns: The execution decision of the launch
	Launch()
	Create()
}

// Responsible for launching a Golang executor. 
// The executor will be launched in the same go routine of the rest of the 
// component, i.e. its driver and publisher.
type InProcessComponentLauncher struct {

}

func (i *InProcessComponentLauncher) Create() {

}

func (i InProcessComponentLauncher) Launch() {

}