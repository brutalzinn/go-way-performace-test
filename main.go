package gowayperformacetest

import (
	"sync"

	"github.com/brutalzinn/go-way-performace-test/mensure"
)

var Version string = "1.0.0"

func MensureFunction(name string, callback func(mensure.TimeBenchmark), fn func(...any), args ...any) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		benchmark := mensure.New(name)
		benchmark.Start()
		fn(args)
		benchmark.End()
		callback(benchmark)
	}()
	wg.Wait()
}
