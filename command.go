package go_system_api

import "time"

type CommandStep struct {
	CommandName string `json:"command_name"`
	Args        string `json:"args"`
	Url         string `json:"url"`
}

type CommandList struct {
	QueryId  string        `json:"query_id"`
	Commands []CommandStep `json:"commands"`
	Step     int           `json:"step"`
	ErrorUrl string        `json:"error_url"`
}

type EventData struct {
	Raw       *string                `json:"raw"`
	IndexTime *time.Time             `json:"index_time"`
	TimeStamp *time.Time             `json:"time_stamp"`
	EventType *string                `json:"event_type"`
	Category  *string                `json:"category"`
	Derived   map[string]interface{} `json:"derived"`
}

type ProcessingEvent struct {
	Commands CommandList `json:"commands"`
	Event    EventData   `json:"event"`
}
