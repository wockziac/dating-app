package dateutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDate(t *testing.T) {
	inputs := []string{"2023-09-10", "2023-09-10 00:00:00", "2023/02/04"}
	for _, dt := range inputs {
		_, err := ParseDate(dt)
		assert.Nil(t, err, fmt.Sprintf("test failed for sample: %v", dt))
	}
}
