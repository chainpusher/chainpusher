package postoffice

type PostOfficeFactory struct {
}

func (p *PostOfficeFactory) CreatePostOffice() PostOffice {
	return &PostOfficeCoroutine{}
}
