package routine

import (
	"github.com/Farengier/gotools/logging"
	"github.com/Farengier/gotools/signals"
	"sync"
	"syscall"
)

var (
	once            sync.Once
	routinesCount   = 0
	isStopRequested = false
	stopChannel     = make(chan struct{})
	exitChannel     = make(chan struct{})
)

//goland:noinspection GoUnusedExportedFunction
func WaitForExit() {
	select {
	case <-exitChannel:
	}
}

//goland:noinspection GoUnusedExportedFunction
func WaitTillShutdownRequested() {
	once.Do(stopOnCtrlC)
	select {
	case <-stopChannel:
	}
}

// Stop requires routines to stop processing
func Stop() {
	if !isStopRequested {
		isStopRequested = true
		close(stopChannel)
	}
}

// StopChannel returns channel to listen. Will be closed if stop required
//goland:noinspection GoUnusedExportedFunction
func StopChannel() <-chan struct{} {
	return stopChannel
}

//goland:noinspection GoUnusedExportedFunction
func IsStopRequested() bool {
	return isStopRequested
}

//goland:noinspection GoUnusedExportedFunction
func RunningRoutines() int {
	return routinesCount
}

//goland:noinspection GoUnusedExportedFunction
func StartRoutine(name string, routine func()) {
	if isStopRequested {
		logging.Warn("Routine ", name, " start declined")
		return
	}
	go func() {
		routineStarted(name)
		defer routineStopped(name)
		routine()
	}()
}

func routineStarted(name string) {
	logging.Debug("Routine ", name, " started")
	routinesCount++
}

func routineStopped(name string) {
	logging.Debug("Routine ", name, " stopped")
	routinesCount--
	if isStopRequested && routinesCount == 0 {
		logging.Info("Exiting")
		close(exitChannel)
	}
}

func stopOnCtrlC() {
	signals.OnSignal(syscall.SIGINT, Stop)
	signals.OnSignal(syscall.SIGTERM, Stop)
}
