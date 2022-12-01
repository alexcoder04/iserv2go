package web

import (
	"encoding/json"
	"fmt"

	"github.com/alexcoder04/iserv2go/iserv/types"
)

func (c *IServWebClient) GetBadges() (map[string]int, error) {
	data, err := c.doGetRequest("/app/navigation/badges")
	if err != nil {
		return map[string]int{}, err
	}

	resData := map[string]int{}
	err = json.Unmarshal(data, &resData)
	return resData, err
}

func (c *IServWebClient) GetNotifications() (*types.NotificationInfo, error) {
	data, err := c.doGetRequest("/user/api/notifications")
	if err != nil {
		return &types.NotificationInfo{}, err
	}

	notInfo := &types.NotificationInfo{}
	err = json.Unmarshal(data, &notInfo)
	return notInfo, err
}

func (c *IServWebClient) GetUpcomingEvents(limit uint) (*types.EventsInfo, error) {
	data, err := c.doGetRequest(fmt.Sprintf("/calendar/api/upcoming?limit=%d&includeSubscriptions=true", limit))
	if err != nil {
		return &types.EventsInfo{}, err
	}

	eventsInfo := &types.EventsInfo{}
	err = json.Unmarshal(data, &eventsInfo)
	return eventsInfo, err
}

func (c *IServWebClient) GetCurrentExercises() ([]types.IServExercise, error) {
	return c.getExercisesFrom("")
}

func (c *IServWebClient) GetPastExercises() ([]types.IServExercise, error) {
	return c.getExercisesFrom("/past/exercise")
}
