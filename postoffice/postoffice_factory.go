package postoffice

type Factory struct {
}

func (p *Factory) CreatePostOffice() PostOffice {
	return &Coroutine{}
}
