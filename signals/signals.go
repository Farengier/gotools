package signals

import (
	"os"
	"os/signal"
	"sync"
)

var (
	once  sync.Once
	calls = make(map[os.Signal][]func())
	ch    = make(chan os.Signal, 1)
)

func OnSignal(sig os.Signal, callback func()) {
	signal.Notify(ch, sig)
	calls[sig] = append(calls[sig], callback)
	once.Do(processor)
}

func processor() {
	go func() {
		for {
			sig := <-ch
			if callbacks, ok := calls[sig]; ok {
				for _, call := range callbacks {
					call()
				}
			}
		}
	}()
}
