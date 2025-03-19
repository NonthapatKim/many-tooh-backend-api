package function

import (
	"encoding/json"
	"fmt"
)

type BoolString bool

func (b *BoolString) UnmarshalJSON(data []byte) error {
	if string(data) == "true" {
		*b = true
		return nil
	} else if string(data) == "false" {
		*b = false
		return nil
	}

	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		if strValue == "true" {
			*b = true
		} else if strValue == "false" {
			*b = false
		} else {
			return fmt.Errorf("invalid value for BoolString: %s", strValue)
		}
		return nil
	}

	return fmt.Errorf("invalid boolean value: %s", data)
}
