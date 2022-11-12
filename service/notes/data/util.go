package data

import (
	"encoding/json"
	"fmt"
)

// Unmarshal loads a SurrealDB response into a struct.
func unmarshal(data, v interface{}) error {
	var ok bool

	assertedData, ok := data.([]interface{})
	if !ok {
		return ErrInvalidResponse
	}
	sliceFlag := isSlice(v)

	var jsonBytes []byte
	var err error
	if !sliceFlag && len(assertedData) > 0 {
		jsonBytes, err = json.Marshal(assertedData[0])
	} else {
		jsonBytes, err = json.Marshal(assertedData)
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBytes, v)
	if err != nil {
		return err
	}

	return err
}

// UnmarshalRaw loads a raw SurrealQL response returned by Query into a struct. Queries that return with results will
// return ok = true, and queries that return with no results will return ok = false.
func unmarshalRaw(rawData, v interface{}) (ok bool, err error) {
	var data []interface{}
	if data, ok = rawData.([]interface{}); !ok {
		return false, ErrInvalidResponse
	}

	var responseObj map[string]interface{}
	if responseObj, ok = data[0].(map[string]interface{}); !ok {
		return false, ErrInvalidResponse
	}

	var status string
	if status, ok = responseObj["status"].(string); !ok {
		return false, ErrInvalidResponse
	}
	if status != statusOK {
		return false, ErrQuery
	}

	result := responseObj["result"]
	if len(result.([]interface{})) == 0 {
		return false, nil
	}
	err = unmarshal(result, v)
	if err != nil {
		return false, err
	}

	return true, nil
}

func isSlice(possibleSlice interface{}) bool {
	slice := false

	switch v := possibleSlice.(type) { //nolint:gocritic
	default:
		res := fmt.Sprintf("%s", v)
		if res == "[]" || res == "&[]" || res == "*[]" {
			slice = true
		}
	}

	return slice
}
