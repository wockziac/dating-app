package phonenumberutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePhoneNumber(t *testing.T) {
	t.Run("incorrect phone number", func(t *testing.T) {
		invalidPhones := []string{
			"+628-221-234-856-7888",
			"081850147",
		}
		for _, number := range invalidPhones {
			_, err := ParsePhoneNumber(number)
			assert.NotNil(t, err, fmt.Sprintf("failed case for sample: %s", number))
		}
	})

	t.Run("incorrect phone number", func(t *testing.T) {
		invalidPhones := map[string]string{
			"+628-221-234-856-7": "6282212348567",
			"0818501471":         "0818501471",
		}
		for sample, expectedNumber := range invalidPhones {
			result, err := ParsePhoneNumber(sample)
			assert.Equal(t, expectedNumber, result, fmt.Sprintf("failed case for sample: %s", sample))
			assert.Nil(t, err)
		}
	})
}

func TestStripPhoneNumber(t *testing.T) {

	t.Run("string phone number symbols", func(t *testing.T) {
		phoneNumSamples := map[string]string{
			"+892-8851-778": "8928851778",
			"8928851778":    "8928851778",
		}

		for sample, expected := range phoneNumSamples {
			result := stripPhoneNumber(sample)
			assert.Equal(t, expected, result, fmt.Sprintf("failed case for sample %s", sample))
		}
	})
}
