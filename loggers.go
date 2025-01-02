/*
####### sdk.loggers (c) 2024 Archivage Numérique ################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package loggers

import (
	"sync"

	"github.com/archnum/sdk.base/logger/manager"
	"github.com/ltrochet/loggers"
)

var (
	_loggers = newLoggers()
)

type (
	Loggers interface {
		All() map[string]manager.Logger
	}

	implLoggers struct {
		ml map[string]manager.Logger
		mu sync.Mutex
	}
)

func newLoggers() *implLoggers {
	return &implLoggers{
		ml: make(map[string]manager.Logger),
	}
}

func (impl *implLoggers) registerLogger(l manager.Logger) {
	impl.mu.Lock()
	defer impl.mu.Unlock()

	impl.ml[l.ID().String()] = l
}

func init() {
	manager.RegisterCallback(_loggers.registerLogger)
}

func (impl *implLoggers) all() map[string]loggers.Logger {
	impl.mu.Lock()
	defer impl.mu.Unlock()

	all := make(map[string]loggers.Logger, len(impl.ml))

	for id, l := range impl.ml {
		all[id] = &implLogger{
			l: l,
		}
	}

	return all
}

func All() map[string]loggers.Logger {
	return _loggers.all()
}

/*
####### END ############################################################################################################
*/
