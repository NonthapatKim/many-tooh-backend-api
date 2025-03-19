package function

import (
	"fmt"
	"time"
)

type CustomDate struct {
	time.Time
}

type CustomDateTime struct {
	time.Time
}

func (c *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1])

	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		return fmt.Errorf("invalid date format: %s", str)
	}
	c.Time = parsedTime
	return nil
}

const timeLayout = "2006-01-02 15:04:05.999999"

func (c *CustomDateTime) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1])

	parsedTime, err := time.Parse(timeLayout, str)
	if err != nil {
		return fmt.Errorf("invalid date format: %s, error: %v", str, err)
	}
	c.Time = parsedTime
	return nil
}
