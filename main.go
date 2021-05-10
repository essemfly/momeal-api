package main

import (
	"log"
	"sync"

	"lessbutter.co/mealkit/crawler"
)

func main() {
	log.Println("Main function Start")
	var wg sync.WaitGroup

	wg.Add(1)
	go crawler.CrawlNaverSearch(&wg, 1)

	wg.Wait()
	log.Println("Main function End")
}
