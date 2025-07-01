package valueobject

import (
	"emailn/internal/shared/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupAll() {
	validator.ConfigValidator()
}

func TestMain(m *testing.M) {
	setupAll()
}

func Test_NewDateTime_Success(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	validDateTime := "2025-04-19T16:23:123"

	/// Act
	dateTime, err := NewDateTime(validDateTime)

	/// Assert
	assert.Nil(err)
	assert.NotNil(dateTime)
	assert.EqualValues(*dateTime, validDateTime)
}

func Test_NewDateTime_InvalidDateTime(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	invalidDateTime := "19/04/2025"

	/// Act
	dateTime, err := NewDateTime(invalidDateTime)

	/// Assert
	assert.Nil(dateTime)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Invalid date time")
}
