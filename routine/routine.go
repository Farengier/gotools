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
)

//goland:noinspection GoUnusedExportedFunction
func WaitTillShutdown() {
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
func StartRoutine(routine func()) {
	go func() {
		routineStarted()
		defer routineStopped()
		routine()
	}()
}

func routineStarted() {
	logging.Info("Started routine")
	routinesCount++
}

func routineStopped() {
	logging.Info("Stopped routine")
	routinesCount--
}

func stopOnCtrlC() {
	signals.OnSignal(syscall.SIGINT, Stop)
	signals.OnSignal(syscall.SIGTERM, Stop)
}
