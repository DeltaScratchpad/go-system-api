package go_system_api

type ErrorBody struct {
	QueryID     string     `json:"query_id"`
	Step        int64      `json:"step"`
	Recoverable bool       `json:"recoverable"`
	ErrorMsg    string     `json:"error_msg"`
	DebugMsg    string     `json:"debug_msg"`
	Event       *EventData `json:"event"`
}
