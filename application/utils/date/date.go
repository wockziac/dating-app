package dateutil

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func ParseDate(date string) (time.Time, error) {
	t, err := dateparse.ParseAny(date)
	if err != nil {
		return t, fmt.Errorf(fmt.Sprintf("invalid input date. err: %v", err))
	}

	return t, nil
}
