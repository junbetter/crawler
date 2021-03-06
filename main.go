package main

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/persist"
	"crawler/schduler"
	"crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		//Scheduler: &schduler.SimpleScheduler{},
		Scheduler:        &schduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	//e.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	//})

	e.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})
}
