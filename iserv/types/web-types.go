package types

type Notification struct{}

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
