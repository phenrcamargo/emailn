package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEmail_Success(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	validEmail := "email@email.com"

	/// Act
	email, err := NewEmail(validEmail)

	/// Assert
	assert.Nil(err)
	assert.NotNil(email)
	assert.EqualValues(*email, validEmail)
}

func Test_NewEmail_Invalid(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	invalidEmail := "email"

	/// Act
	email, err := NewEmail(invalidEmail)

	/// Assert
	assert.Nil(email)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Invalid email")
}
