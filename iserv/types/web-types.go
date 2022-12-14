package types

import (
	"time"
)

// exercise
type Exercise struct {
	Description string
	DueDate     time.Time
	Files       []string
	Id          int
	StartDate   time.Time
	Tags        []string
	Teacher     string
	Title       string
	Types       []string
}

// notification
type Notification struct {
	Type         string     `json:"type"`
	Id           int        `json:"id"`
	GroupId      string     `json:"groupId"`
	GroupTitle   string     `json:"groupTitle"`
	AutoGrouping bool       `json:"autoGrouping"`
	Message      string     `json:"message"`
	GroupMessage string     `json:"groupMessage"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	Trigger      string     `json:"trigger"`
	Url          string     `json:"url"`
	Icon         string     `json:"icon"`
	Date         IServTime1 `json:"date"`
}

type NotificationData struct {
	LastEventId   int            `json:"lastEventId"`
	LastId        int            `json:"lastId"`
	Since         int            `json:"since"`
	Count         int            `json:"count"`
	Notifications []Notification `json:"notifications"`
}

type NotificationInfo struct {
	Status string           `json:"status"`
	Data   NotificationData `json:"data"`
}

// event
type Event struct {
	Id                     string     `json:"id"`
	Uid                    string     `json:"uid"`
	RecurrenceId           string     `json:"recurrenceId"`
	Hash                   string     `json:"hash"`
	Title                  string     `json:"title"`
	Description            string     `json:"description"`
	DescriptionHtml        string     `json:"descriptionHtml"`
	Category               string     `json:"category"`
	CategoryColor          string     `json:"category_color"`
	Start                  IServTime1 `json:"start"`
	End                    IServTime1 `json:"end"`
	Timezone               string     `json:"timezone"`
	Editable               bool       `json:"editable"`
	Deletable              bool       `json:"deletable"`
	AllDay                 bool       `json:"allDay"`
	Recurring              bool       `json:"recurring"`
	CalendarId             string     `json:"calendarId"`
	Location               string     `json:"location"`
	LocationHtml           string     `json:"locationHtml"`
	When                   string     `json:"when"`
	Organizer              string     `json:"organizer"`
	Creator                string     `json:"creator"`
	CreatedAt              IServTime2 `json:"createdAt"`
	Status                 string     `json:"status"`
	ParticipantsWithStatus string     `json:"participantsWithStatus"`
	CurrentUserPartstat    string     `json:"currentUserPartstat"`
	ShowAttendanceButtons  bool       `json:"showAttendanceButtons"`
	IsOrganizer            bool       `json:"isOrganizer"`
	Color                  string     `json:"color"`
	CalendarName           string     `json:"calendarName"`
	Alarms                 []Alarm    `json:"alarms"`
}

type Alarm struct {
	Action            string   `json:"action"`
	Trigger           string   `json:"trigger"`
	TriggerParameters []string `json:"trigger_parameters"`
}

type EventsInfo struct {
	Events []Event  `json:"events"`
	Errors []string `json:"errors"`
}
