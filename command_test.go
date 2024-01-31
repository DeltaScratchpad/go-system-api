package go_system_api

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	// Create an example ProcessingEvent
	var eventRaw string = "Test Raw"
	timeNow := time.Now()
	derived := make(map[string]interface{})
	derived["test"] = "test"
	derived["test2"] = 40
	derived["test3"] = make(map[string]interface{})
	derived["test3"].(map[string]interface{})["test4"] = "test4"

	var pe = ProcessingEvent{
		Commands: CommandList{
			QueryId: "Test Command 1",
			Commands: []CommandStep{
				{CommandName: "Command 1", Args: "args 1", Url: "Url 1"},
				{CommandName: "Command 2", Args: "args 2", Url: "Url 2"},
				{CommandName: "Command 3", Args: "args 3", Url: "Url 3"},
			},
			Step:     0,
			ErrorUrl: "Error Url",
		},
		Event: EventData{
			Raw:       &eventRaw,
			IndexTime: &timeNow,
			TimeStamp: &timeNow,
			EventType: nil,
			Category:  nil,
			Derived:   &derived,
		},
	}

	jsonData, err := json.Marshal(pe)
	if err != nil {
		t.Errorf("Error marshalling ProcessingEvent: %s", err)
	}
	println(string(jsonData))

	// Parse the jsonData back into a ProcessingEvent
	var pe2 ProcessingEvent
	err = json.Unmarshal(jsonData, &pe2)
	if err != nil {
		t.Errorf("Error unmarshalling ProcessingEvent: %s", err)
	}
}
