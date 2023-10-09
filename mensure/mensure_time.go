package mensure

import (
	"fmt"
	"time"
)

type TimeBenchmark struct {
	Name      string
	StartTime time.Time
	ElapsedMS int64
	EndTime   time.Time
}

func New(name string) TimeBenchmark {
	timeBenchMark := TimeBenchmark{
		Name: name,
	}
	return timeBenchMark
}

func (timeBenchMark *TimeBenchmark) ChangeName(name string) {
	timeBenchMark.Name = name
}

func (timeBenchMark *TimeBenchmark) Start() {
	timeBenchMark.StartTime = time.Now()
}

func (timeBenchMark *TimeBenchmark) End() {
	timeBenchMark.EndTime = time.Now()
	elapsedMS := timeBenchMark.EndTime.Sub(timeBenchMark.StartTime).Milliseconds()
	timeBenchMark.ElapsedMS = elapsedMS
}

func (timeBenchMark *TimeBenchmark) ToString() string {
	toTxt := fmt.Sprintf("%s FINISHED in %dms", timeBenchMark.Name, timeBenchMark.ElapsedMS)
	return toTxt
}
