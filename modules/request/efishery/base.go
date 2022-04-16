package efishery

type req struct{}

func NewRequest() Interface {
	return &req{}
}
