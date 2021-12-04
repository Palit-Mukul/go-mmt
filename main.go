package main

import (
	"encoding/json"
	"fmt"
	"github.com/Palit-Mukul/go-mmt/service"
	"log"
	"sync"
)

var testConfig = `{
		"config": [
			{
       			"url":"https://asia-east2-jsondoc.cloudfunctions.net/function-1?delay=1000",
       			"isParallel": true,
       			"count": "3"
			},
			{
       			"url":"https://asia-east2-jsondoc.cloudfunctions.net/function-1?delay=10000", 
       			"isParallel": false,
       			"count": "1"
			},
			{
       			"url":"https://asia-east2-jsondoc.cloudfunctions.net/function-1?delay=2000", 
       			"isParallel": true,
       			"count": "3"
			},
			{
       			"url":"https://asia-east2-jsondoc.cloudfunctions.net/function-1?delay=1000",
       			"isParallel": false,
       			"count": "2"
			}
		]
}`

func main() {
	var mainWG sync.WaitGroup
	parsedConfig := ParseJSON()
	for _,v :=range parsedConfig.Config {
		mainWG.Add(1)
		go service.NewService(&mainWG, v)
	}
	fmt.Println("main thread unblocked")
	mainWG.Wait()
	fmt.Println("Total Time : ",service.TotalTime)
}

func ParseJSON () service.ProgramConfig {
	var programConfig service.ProgramConfig
	err := json.Unmarshal([]byte(testConfig), &programConfig)
	if err != nil {
		log.Fatalf("could not unmarshal json : %s", err)
	}
	return programConfig
}

















