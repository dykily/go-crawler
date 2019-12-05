package main

import (
	"./zhenai/parser"
	"go-crawler/engine"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,

	})
}

