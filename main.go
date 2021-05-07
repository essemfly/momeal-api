package main

import (
	"log"
	"sync"

	"lessbutter.co/mealkit/crawler"
)

func main() {
	log.Println("Main function Start")
	var wg sync.WaitGroup

	var pages = []string{"1", "2"}
	for _, pageNum := range pages {
		wg.Add(1)
		go crawler.CrawlFreshEasy(&wg, pageNum)
	}
	wg.Wait()
	log.Println("Main function End")
}
