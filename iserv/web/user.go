package web

import "encoding/json"

func (c *IServWebClient) GetBadges() (map[string]int, error) {
	data, err := c.DoGetRequest("/app/navigation/badges")
	if err != nil {
		return map[string]int{}, err
	}

	resData := map[string]int{}
	err = json.Unmarshal(data, &resData)
	return resData, err
}

//func (c *IServWebClient) GetExercises() ([]IServExercise, error) {
//	res, err := c.HttpClient.Get(c.Config.IServURL + "app/navigation/badges")
//	if err != nil {
//		return []IServExercise{}, err
//	}
//	defer res.Body.Close()
//
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		return []IServExercise{}, err
//	}
//	tasksTable := doc.Find("#crud-table")
//	fmt.Println(tasksTable)
//	for _, tr := range tasksTable.Filter("tr").Nodes {
//		fmt.Println(tr)
//	}
//	return []IServExercise{}, nil
//}
