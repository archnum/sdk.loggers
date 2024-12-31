/*
####### sdk.loggers (c) 2024 Archivage Numérique ################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package loggers

import (
	"fmt"

	"github.com/archnum/sdk.base/logger/manager"
)

type (
	implLogger struct {
		l manager.Logger
	}
)

func (impl *implLogger) ID() string {
	return string(impl.l.ID())
}

func (impl *implLogger) Name() string {
	return impl.l.Name()
}

func (impl *implLogger) Level() string {
	return impl.l.Level().String()
}

func (impl *implLogger) DecLog() {
	// TODO
	fmt.Println("DEC", impl.l.ID())
}

func (impl *implLogger) IncLog() {
	// TODO
	fmt.Println("INC", impl.l.ID())
}

/*
####### END ############################################################################################################
*/
