package helpers

import (
	"brutalzinn/go-way-performace-test/mensure"
	"log"
	"os"
)

type LoggerMetric struct {
	OutputLogFile string
}

func (lm *LoggerMetric) Write(data any) {
	f, err := os.OpenFile(lm.OutputLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(data)
}

func (lm *LoggerMetric) WriteMensureTime(data mensure.TimeBenchmark) {
	lm.Write(data.ToString())
}
