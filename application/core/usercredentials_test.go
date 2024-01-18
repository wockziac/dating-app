package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserCredentials(t *testing.T) {

	t.Run("empty email address", func(t *testing.T) {
		uc := UserCredentials{}
		err := uc.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "empty email address")
	})

	t.Run("invalid email address", func(t *testing.T) {
		invalidEmails := []string{
			"waingapuw@gmailcom",
			"123.com",
			"#waingapuw@gmail.com",
			"waingapuw@com",
		}

		uc := UserCredentials{}
		for _, email := range invalidEmails {
			uc.EmailAddress = email
			err := uc.Validate()
			assert.NotNil(t, err)
		}
	})

	t.Run("empty phone number", func(t *testing.T) {
		uc := UserCredentials{
			EmailAddress: "waingapu@gmail.com",
		}
		err := uc.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "empty phone number")
	})

	t.Run("invalid phone number", func(t *testing.T) {
		invalidPhones := []string{
			"+628-221-234-856-7888",
			"081850147",
		}
		uc := UserCredentials{
			EmailAddress: "waingapu@gmail.com",
		}
		for _, number := range invalidPhones {
			uc.PhoneNumber = number
			err := uc.Validate()
			assert.NotNil(t, err, fmt.Sprintf("failed case for sample: %s", number))
		}
	})
}
