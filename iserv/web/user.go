package web

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alexcoder04/iserv2go/iserv/types"
)

func (c *WebClient) GetBadges() (map[string]int, error) {
	data, err := c.doGetRequest("/app/navigation/badges")
	if err != nil {
		return map[string]int{}, err
	}

	if strings.TrimSpace(string(data)) == "[]" {
		return map[string]int{}, nil
	}

	resData := map[string]int{}
	err = json.Unmarshal(data, &resData)
	return resData, err
}

func (c *WebClient) GetNotifications() (*types.NotificationInfo, error) {
	data, err := c.doGetRequest("/user/api/notifications")
	if err != nil {
		return &types.NotificationInfo{}, err
	}

	notInfo := &types.NotificationInfo{}
	err = json.Unmarshal(data, &notInfo)
	return notInfo, err
}

func (c *WebClient) GetUpcomingEvents(limit uint) (*types.EventsInfo, error) {
	data, err := c.doGetRequest(fmt.Sprintf("/calendar/api/upcoming?limit=%d&includeSubscriptions=true", limit))
	if err != nil {
		return &types.EventsInfo{}, err
	}

	eventsInfo := &types.EventsInfo{}
	err = json.Unmarshal(data, &eventsInfo)
	return eventsInfo, err
}

func (c *WebClient) GetCurrentExercises() ([]types.Exercise, error) {
	return c.getExercisesFrom("")
}

func (c *WebClient) GetPastExercises() ([]types.Exercise, error) {
	return c.getExercisesFrom("/past/exercise")
}
