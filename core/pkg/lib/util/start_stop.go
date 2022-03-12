package util

type Starter interface {
	Start()
}

type Stopper interface {
	Stop()
}

type StartStopper interface {
	Starter
	Stopper
}
