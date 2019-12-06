package main

import (
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,

	})
}

