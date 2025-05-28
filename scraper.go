package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/pratimeshtiwari/rssaggregator/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scrapping started with concurrency: %d, time between requests: %s", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		courses, err := db.GetNextCoursesToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fetching courses to scrape: %v", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, course := range courses {
			wg.Add(1)
			go scrapeFeed(db, wg, course)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, course database.Course) {
	defer wg.Done()
	_, error := db.MarkCourseAsFetched(context.Background(), course.ID)
	if error != nil {
		log.Printf("Error marking course as fetched: %v", error)
		return
	}
	rssCourse, err := urlToCourse(course.Url)
	if err != nil {
		log.Printf("Error fetching course from URL %s: %v", course.Url, err)
		return
	}
	for _, item := range rssCourse.Channel.Item {
		log.Println("Found course : ", item.Title)
	}
	log.Printf("Course %s fetched , %v items found", course.ID, len(rssCourse.Channel.Item))
}
