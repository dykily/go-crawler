package main

import (
	"go-crawler/engine"
	"go-crawler/persist"
	"go-crawler/scheduler"
	"go-crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,

	})
}

