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
		all   map[string]manager.Logger
		mutex sync.Mutex
	}
)

func newLoggers() *implLoggers {
	return &implLoggers{
		all: make(map[string]manager.Logger),
	}
}

func (impl *implLoggers) registerLogger(l manager.Logger) {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()

	impl.all[string(l.ID())] = l
}

func init() {
	manager.RegisterCallback(_loggers.registerLogger)
}

func All() map[string]loggers.Logger {
	return nil
}

/*
####### END ############################################################################################################
*/
