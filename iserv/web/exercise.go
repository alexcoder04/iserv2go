package web

import (
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexcoder04/iserv2go/iserv/types"
)

func getExerciseInfo(url string) (types.IServExercise, error) {
	return types.IServExercise{}, nil
}

func (c *IServWebClient) getExercisesFrom(path string) ([]types.IServExercise, error) {
	// get html page of exercise overview
	doc, err := c.doGetRequestQueryDoc("/exercise" + path)
	if err != nil {
		return []types.IServExercise{}, err
	}

	// for each task get the url
	var urls []string
	doc.Find("#crud-table tbody tr").Each(func(i int, s *goquery.Selection) {
		if !s.HasClass("info") {
			url, _ := s.Children().Eq(1).Children().First().Attr("href")
			urls = append(urls, url)
		}
	})

	// get info for all exercises concurrently
	exercises := []types.IServExercise{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)

		go func(url string, wg *sync.WaitGroup) {
			exercise, err := getExerciseInfo(url)
			if err != nil {
				return
			}

			mu.Lock()
			exercises = append(exercises, exercise)
			mu.Unlock()

			wg.Done()
		}(url, &wg)
	}

	wg.Wait()

	return exercises, nil
}
