package routine

import (
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
func StopChannel() <-chan struct{} {
	return stopChannel
}

func IsStopRequested() bool {
	return isStopRequested
}

func RunningRoutines() int {
	return routinesCount
}

func StartRoutine(routine func()) {
	go func() {
		routineStarted()
		defer routineStopped()
		routine()
	}()
}

func routineStarted() {
	routinesCount++
}

func routineStopped() {
	routinesCount--
}

func stopOnCtrlC() {
	signals.OnSignal(syscall.SIGINT, Stop)
	signals.OnSignal(syscall.SIGTERM, Stop)
}
