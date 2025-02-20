package application

import "context"

type AppContextBase struct {
	Context context.Context
}

type AppRunner interface {
	Start() (bool, error)
	Stop() (bool, error)
	Init() (bool, error)
}
