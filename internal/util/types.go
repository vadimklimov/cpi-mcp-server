package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	if t.IsZero() {
		return ""
	}

	return t.Format(time.RFC3339)
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	parsers := []func(string) (Timestamp, error){
		parseEpochTimestamp,
		parseDateTimeTimestamp,
	}

	var input string
	if err := json.Unmarshal(data, &input); err != nil {
		return fmt.Errorf("failed to unmarshal timestamp: %w", err)
	}

	for _, parser := range parsers {
		timestamp, err := parser(input)
		if err == nil {
			*t = timestamp

			return nil
		}
	}

	return fmt.Errorf("failed to parse timestamp")
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}

	output, err := json.Marshal(t.Format(time.RFC3339))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal timestamp: %w", err)
	}

	return output, nil
}

func parseEpochTimestamp(input string) (Timestamp, error) {
	timestamp, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return Timestamp{}, fmt.Errorf("failed to parse value as epoch timestamp: %w", err)
	}

	output := Timestamp{time.UnixMilli(timestamp).UTC()}

	return output, nil
}

func parseDateTimeTimestamp(input string) (Timestamp, error) {
	timestamp, err := time.Parse("2006-01-02T15:04:05.000", input)
	if err != nil {
		return Timestamp{}, fmt.Errorf("failed to parse value as datetime timestamp: %w", err)
	}

	output := Timestamp{timestamp.UTC()}

	return output, nil
}
