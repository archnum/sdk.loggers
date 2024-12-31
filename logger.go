/*
####### sdk.loggers (c) 2024 Archivage Numérique ################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package loggers

import (
	"github.com/archnum/sdk.base/logger/level"
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
	lvl := impl.l.Level()
	if lvl == level.Error {
		return
	}

	impl.l.SetLevel(level.Level(lvl + 1))
}

func (impl *implLogger) IncLog() {
	lvl := impl.l.Level()
	if lvl == level.Trace {
		return
	}

	impl.l.SetLevel(level.Level(lvl - 1))
}

/*
####### END ############################################################################################################
*/
