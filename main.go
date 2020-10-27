package main

import (
	"syscall/js"
	"time"
)

func timeString() string {
	return time.Now().Format(time.ANSIC)
}

func CurrentTime(this js.Value, inputs []js.Value) interface{} {
	return timeString()
}

func EverySecond(this js.Value, inputs []js.Value) interface{} {
	tick := time.Tick(time.Second)
	callback := inputs[len(inputs)-1:][0]
	go func() {
		for range tick {
			callback.Invoke(timeString())
		}
	}()

	return nil
}

func main() {
	c := make(chan bool)
	js.Global().Set("currentTime", js.FuncOf(CurrentTime))
	js.Global().Set("everySecond", js.FuncOf(EverySecond))
	<-c
}
