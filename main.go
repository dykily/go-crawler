package main

import (
	"go-crawler/scheduler"
	"go-crawler/zhenai/parser"
	"go-crawler/engine"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,

	})
}

