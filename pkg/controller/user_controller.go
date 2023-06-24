package controller

// Controllers should receive all adapters needed to do its work. Logger and Metrics too.
type UserController struct{}

func NewUserController() *UserController {
	return new(UserController)
}

func (c *UserController) GetUser() string {
	return "Hello World"
}
