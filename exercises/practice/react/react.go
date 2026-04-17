package react

// Define types that will provide the Reactor, cell, and Canceler interface here.
// The prepopulated code assumes types
// canceler <- Canceler
// myCell <- cell, ComputeCell, InputCell
// reactor <- Reactor

func (c *canceler) Cancel() {
	panic("Please implement the Cancel function")
}

func (c *myCell) Value() int {
	panic("Please implement the Value function")
}

func (c *myCell) SetValue(value int) {
	panic("Please implement the SetValue function")
}

func (c *myCell) AddCallback(callback func(int)) Canceler {
	panic("Please implement the AddCallback function")
}

func New() Reactor {
	panic("Please implement the New function")
}

func (r *reactor) CreateInput(initial int) InputCell {
	panic("Please implement the CreateInput function")
}

func (r *reactor) CreateCompute1(dep cell, compute func(int) int) ComputeCell {
	panic("Please implement the CreateCompute1 function")
}

func (r *reactor) CreateCompute2(dep1, dep2 cell, compute func(int, int) int) ComputeCell {
	panic("Please implement the CreateCompute2 function")
}
