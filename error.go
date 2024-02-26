package go_system_api

import "encoding/json"

type ErrorBody struct {
	QueryID     string     `json:"query_id"`
	Step        int64      `json:"step"`
	Recoverable bool       `json:"recoverable"`
	ErrorMsg    string     `json:"error_msg"`
	DebugMsg    string     `json:"debug_msg"`
	Event       *EventData `json:"event"`
}

func (e ErrorBody) String() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		panic("Failed to serialise.")
	}
	return string(bytes)
}

func (e ErrorBody) Error() string {
	return e.ErrorMsg
}
