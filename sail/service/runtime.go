package service

import "github.com/sourcegraph/conc"

type Runtime interface {
	Start(wg *conc.WaitGroup)
	Stop()
}
