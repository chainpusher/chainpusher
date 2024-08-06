package facade

type IntentServiceFacade interface {
	CreateIntent() (string, error)
}
