package converter

type req struct{}

func NewRequest() Interface {
	return &req{}
}
