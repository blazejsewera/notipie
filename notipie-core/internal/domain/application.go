package domain

type Application struct {
	Name         string
	SmallIconURL string
	BigIconURL   string
	Listeners    map[string]*Listener
}

func NewApplication(name string, smallIconURL string, bigIconURL string) *Application {
	return &Application{
		Name:         name,
		SmallIconURL: smallIconURL,
		BigIconURL:   bigIconURL,
		Listeners:    make(map[string]*Listener),
	}
}

func (a *Application) RegisterListener(ID string, listener *Listener) {
	a.Listeners[ID] = listener
}

func (a *Application) DeregisterListener(ID string) {
	delete(a.Listeners, ID)
}
