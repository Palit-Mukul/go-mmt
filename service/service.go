package service

import (
	"fmt"
	"log"
"net/http"
"strconv"
"sync"
"time"
)

var TotalTime int

type ProgramConfig struct {
	Config []ConfigJSON `json:"config"`
}

type ConfigJSON struct {
	URL        string `json:"url"`
	IsParallel bool   `json:"isParallel"`
	Count      string `json:"count"`
}

func NewService (timeI chan int, v ConfigJSON)  {
	var wg sync.WaitGroup
	var err error
	count := 0
	count, err = strconv.Atoi(v.Count)
	if err != nil {
		log.Fatalf("could not convert count from string to int : %s", err)
	}
	switch v.IsParallel {
	case true :
		wg.Add(count)
		startTime := time.Now()
		for i:=0;i<count;i++ {
			go HitURLinParallel(&wg,v.URL)
		}
		wg.Wait()
		t := int(time.Since(startTime).Seconds())
		fmt.Println(t)
		timeI <- t
	case false:
		startTime := time.Now()
		for i:=0;i<count;i++ {
			HitURL(v.URL)
		}
		t := int(time.Since(startTime).Seconds())
		fmt.Println(t)
		timeI <- t
		TotalTime += t
	}
}

func HitURL (url string) {
	log.Println("hitting URL : ",url)
	_, err := http.Get(url)
	if err != nil {
		log.Fatalf("could not hit url [%s] : %s",url,err)
	}
}

func HitURLinParallel (wg *sync.WaitGroup, url string) {
	defer wg.Done()
	log.Println("Hitting URL : ", url)
	_, err := http.Get(url)
	if err != nil {
		log.Fatalf("could not hit url [%s] : %s", url, err)
	}
}