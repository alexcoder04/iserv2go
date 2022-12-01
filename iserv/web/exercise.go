package web

import (
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexcoder04/iserv2go/iserv/types"
)

func (c *WebClient) getExerciseInfo(url string) (types.Exercise, error) {
	doc, err := c.doGetRequestQueryDoc(url)
	if err != nil {
		return types.Exercise{}, err
	}

	exercise := types.Exercise{}

	// title
	exercise.Title = doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div.panel-heading > h3").Text()

	// all types of submitting an answer
	exercise.Types = []string{}
	if doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(3) > div > div.panel-body > div.row.pb-3 > form > div > h5").Text() == "Text" {
		exercise.Types = append(exercise.Types, "Text")
	}

	if doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(3) > div > div.panel-body > div.form-group.p-3.m-0.confirmation-flow.confirmation-warning > label").Text() == "Erledigt" {
		exercise.Types = append(exercise.Types, "Mark")
	}

	// dates
	duedate := doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(2) > div > table > tbody > tr > td:nth-child(3) > ul > li:nth-child(1)").Text()
	startdate := doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(2) > div > table > tbody > tr > td:nth-child(2)").Text()
	exercise.DueDate, _ = time.Parse("02.01.2006 15:04", duedate)
	exercise.StartDate, _ = time.Parse("02.01.2006 15:04", startdate)

	// description
	exercise.Description = doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(2) > div > div").Text()

	// get teacher
	exercise.Teacher = doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(2) > div > table > tbody > tr > td.bt0.pt-0.pl-0 > a").Text()

	// files
	exercise.Files = []string{}
	doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(3) > div > form > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		fileurl, _ := s.Children().Eq(1).Children().First().Attr("href")
		exercise.Files = append(exercise.Files, "https://iserv-schillerschule.de/"+fileurl)
	})

	return exercise, nil
}

func (c *WebClient) getExercisesFrom(path string) ([]types.Exercise, error) {
	// get html page of exercise overview
	doc, err := c.doGetRequestQueryDoc("/exercise" + path)
	if err != nil {
		return []types.Exercise{}, err
	}

	// for each task get the url
	var urls []string
	doc.Find("#crud-table tbody tr").Each(func(i int, s *goquery.Selection) {
		if !s.HasClass("info") {
			url, _ := s.Children().Eq(1).Children().First().Attr("href")
			urls = append(urls, strings.TrimLeft(url, c.iServUrl))
		}
	})

	// get info for all exercises concurrently
	exercises := []types.Exercise{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)

		go func(url string, wg *sync.WaitGroup) {
			exercise, err := c.getExerciseInfo(url)
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
