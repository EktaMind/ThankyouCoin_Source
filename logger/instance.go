package logger

import (
	"github.com/EktaMind/Thank_ethereum/log"
)

type Instance struct {
	Log log.Logger
}

func MakeInstance() Instance {
	return Instance{
		Log: log.New(),
	}

}

func (i *Instance) SetName(name string) {
	i.Log = log.New("name", name)
}
