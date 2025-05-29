package model

import (
	"encoding/json"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/test/serdesertest/util"
)

func TestSerDserSaveScheduleRequest(t *testing.T) {
	// 1. Load JSON template
	jsonTemplate, err := util.GetJsonString("SaveScheduleRequest")
	if err != nil {
		t.Fatalf("Failed to load JSON template: %v", err)
	}

	// 2. Deserialize JSON to struct
	var saveScheduleRequest model.SaveScheduleRequest
	err = json.Unmarshal([]byte(jsonTemplate), &saveScheduleRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON: %v", err)
	}

	// 3. Validate deserialized fields
	// String fields
	if saveScheduleRequest.CreatedBy != "sample_createdBy" {
		t.Errorf("Expected CreatedBy = 'sample_createdBy', got '%s'", saveScheduleRequest.CreatedBy)
	}
	if saveScheduleRequest.CronExpression != "sample_cronExpression" {
		t.Errorf("Expected CronExpression = 'sample_cronExpression', got '%s'", saveScheduleRequest.CronExpression)
	}
	if saveScheduleRequest.Description != "sample_description" {
		t.Errorf("Expected Description = 'sample_description', got '%s'", saveScheduleRequest.Description)
	}
	if saveScheduleRequest.Name != "sample_name" {
		t.Errorf("Expected Name = 'sample_name', got '%s'", saveScheduleRequest.Name)
	}
	if saveScheduleRequest.UpdatedBy != "sample_updatedBy" {
		t.Errorf("Expected UpdatedBy = 'sample_updatedBy', got '%s'", saveScheduleRequest.UpdatedBy)
	}
	if saveScheduleRequest.ZoneId != "sample_zoneId" {
		t.Errorf("Expected ZoneId = 'sample_zoneId', got '%s'", saveScheduleRequest.ZoneId)
	}

	// Boolean fields
	if !saveScheduleRequest.Paused {
		t.Errorf("Expected Paused = true, got %v", saveScheduleRequest.Paused)
	}
	if !saveScheduleRequest.RunCatchupScheduleInstances {
		t.Errorf("Expected RunCatchupScheduleInstances = true, got %v", saveScheduleRequest.RunCatchupScheduleInstances)
	}

	// Int64 fields
	if saveScheduleRequest.ScheduleEndTime != 123 {
		t.Errorf("Expected ScheduleEndTime = 123, got %d", saveScheduleRequest.ScheduleEndTime)
	}
	if saveScheduleRequest.ScheduleStartTime != 123 {
		t.Errorf("Expected ScheduleStartTime = 123, got %d", saveScheduleRequest.ScheduleStartTime)
	}

	// Pointer to struct field
	if saveScheduleRequest.StartWorkflowRequest == nil {
		t.Errorf("StartWorkflowRequest should not be nil")
	} else {
		// Validate nested StartWorkflowRequest fields
		if saveScheduleRequest.StartWorkflowRequest.Name != "sample_name" {
			t.Errorf("Expected StartWorkflowRequest.Name = 'sample_name', got '%s'", saveScheduleRequest.StartWorkflowRequest.Name)
		}
		if saveScheduleRequest.StartWorkflowRequest.Version != 123 {
			t.Errorf("Expected StartWorkflowRequest.Version = 123, got %d", saveScheduleRequest.StartWorkflowRequest.Version)
		}
		if saveScheduleRequest.StartWorkflowRequest.Priority != 123 {
			t.Errorf("Expected StartWorkflowRequest.Priority = 123, got %d", saveScheduleRequest.StartWorkflowRequest.Priority)
		}
		if saveScheduleRequest.StartWorkflowRequest.CorrelationId != "sample_correlationId" {
			t.Errorf("Expected StartWorkflowRequest.CorrelationId = 'sample_correlationId', got '%s'", saveScheduleRequest.StartWorkflowRequest.CorrelationId)
		}
		if saveScheduleRequest.StartWorkflowRequest.CreatedBy != "sample_createdBy" {
			t.Errorf("Expected StartWorkflowRequest.CreatedBy = 'sample_createdBy', got '%s'", saveScheduleRequest.StartWorkflowRequest.CreatedBy)
		}

		// Check nested maps
		input := saveScheduleRequest.StartWorkflowRequest.Input
		if input == nil {
			t.Errorf("StartWorkflowRequest.Input should not be nil")
		} else if m, ok := input.(map[string]interface{}); !ok {
			t.Errorf("StartWorkflowRequest.Input should be of type map[string]interface{}")
		} else if len(m) == 0 {
			t.Errorf("StartWorkflowRequest.Input map should not be empty")
		}
		if saveScheduleRequest.StartWorkflowRequest.TaskToDomain == nil || len(saveScheduleRequest.StartWorkflowRequest.TaskToDomain) == 0 {
			t.Errorf("StartWorkflowRequest.TaskToDomain map should not be nil or empty")
		}
	}

	// 4. Serialize struct back to JSON
	serializedJSON, err := json.Marshal(saveScheduleRequest)
	if err != nil {
		t.Fatalf("Failed to serialize struct: %v", err)
	}

	// 5. Round-trip integrity check
	// Serialize and deserialize again to compare structs instead of JSON maps
	var roundTripRequest model.SaveScheduleRequest
	err = json.Unmarshal(serializedJSON, &roundTripRequest)
	if err != nil {
		t.Fatalf("Failed to deserialize round-trip JSON: %v", err)
	}

	// Compare key fields to ensure data integrity
	if roundTripRequest.Name != saveScheduleRequest.Name {
		t.Errorf("Round-trip Name mismatch: %s vs %s", saveScheduleRequest.Name, roundTripRequest.Name)
	}
	if roundTripRequest.CronExpression != saveScheduleRequest.CronExpression {
		t.Errorf("Round-trip CronExpression mismatch: %s vs %s", saveScheduleRequest.CronExpression, roundTripRequest.CronExpression)
	}
	if roundTripRequest.Paused != saveScheduleRequest.Paused {
		t.Errorf("Round-trip Paused mismatch: %v vs %v", saveScheduleRequest.Paused, roundTripRequest.Paused)
	}
	if roundTripRequest.ScheduleStartTime != saveScheduleRequest.ScheduleStartTime {
		t.Errorf("Round-trip ScheduleStartTime mismatch: %d vs %d", saveScheduleRequest.ScheduleStartTime, roundTripRequest.ScheduleStartTime)
	}

	// Compare nested StartWorkflowRequest key fields
	if (saveScheduleRequest.StartWorkflowRequest == nil) != (roundTripRequest.StartWorkflowRequest == nil) {
		t.Errorf("Round-trip StartWorkflowRequest nil mismatch")
	} else if saveScheduleRequest.StartWorkflowRequest != nil {
		if roundTripRequest.StartWorkflowRequest.Name != saveScheduleRequest.StartWorkflowRequest.Name {
			t.Errorf("Round-trip StartWorkflowRequest.Name mismatch: %s vs %s",
				saveScheduleRequest.StartWorkflowRequest.Name, roundTripRequest.StartWorkflowRequest.Name)
		}
	}
}
