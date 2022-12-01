package web

import (
	"time"
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

func (c *IServWebClient) getExerciseInfo(url string) (types.IServExercise, error) {
	// get html page of specific exercise
	res, err := c.httpClient.Get(url)
	if err != nil {
		return types.IServExercise{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return types.IServExercise{}, err
	}

	// parse html page
	exercise := types.IServExercise{} // create struct
	// get title
	title := doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div.panel-heading > h3").Text()
	exercise.Title = title

	// get all types of submitting an answer
	var types []string
	if doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(3) > div > div.panel-body > div.row.pb-3 > form > div > h5").Text() == "Text" {
		types = append(types, "Text")
	}

	if doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(3) > div > div.panel-body > div.form-group.p-3.m-0.confirmation-flow.confirmation-warning > label").Text() == "Erledigt" {
		types = append(types, "Mark")
	}

	exercise.Types = types

	// get dates
	duedate := doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(2) > div > table > tbody > tr > td:nth-child(3) > ul > li:nth-child(1)").Text()
	startdate := doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(2) > div > table > tbody > tr > td:nth-child(2)").Text()
	exercise.DueDate, _ = time.Parse("02.01.2006 15:04", duedate)
	exercise.StartDate, _ = time.Parse("02.01.2006 15:04", startdate)

	// get description
	description := doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(2) > div > div").Text()
	exercise.Description = description

	// get teacher
	teacher := doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(2) > div > table > tbody > tr > td.bt0.pt-0.pl-0 > a").Text()
	exercise.Teacher = teacher

	// files
	var fileurls []string
	doc.Find("#content > div:nth-child(2) > div > div > div:nth-child(2) > div > div:nth-child(3) > div > form > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		fileurl, _ := s.Children().Eq(1).Children().First().Attr("href")
		fileurls = append(fileurls, "https://iserv-schillerschule.de/"+fileurl)
	})
	exercise.Files = fileurls

	// print out
	// fmt.Println(title)
	// fmt.Println(duedate)
	// fmt.Println(startdate)
	// fmt.Println(teacher)
	// fmt.Println("Fileurls:")
	// for _, fileurl := range fileurls {
	// 	fmt.Println(fileurl)
	// }
	// fmt.Println("Types:")
	// for _, submittype := range types {
	// 	fmt.Println(submittype)
	// }
	// fmt.Println(description)

	return exercise, nil
}
