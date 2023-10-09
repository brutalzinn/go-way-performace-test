package main

import (
	"brutalzinn/go-way-performace-test/helpers"
	"brutalzinn/go-way-performace-test/mensure"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	logger := helpers.LoggerMetric{
		OutputLogFile: "test_result/test.txt",
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		benchmark := mensure.New("test doSomething()")
		benchmark.Start()
		time.Sleep(2 * time.Second)
		benchmark.End()
		logger.WriteMensureTime(benchmark)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		benchmark := mensure.New("test doSomething2()")
		benchmark.Start()
		logger.Write("simulate doSomething2()")
		time.Sleep(1 * time.Second)
		logger.Write("end simulate doSomething2()")
		benchmark.End()
		logger.WriteMensureTime(benchmark)
	}()
	wg.Wait()
}
