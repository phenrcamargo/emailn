package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewName_Success(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	validName := "Pedro"

	/// Act
	name, err := NewName(validName)

	/// Assert
	assert.Nil(err)
	assert.NotNil(name)
	assert.EqualValues(*name, validName)
}

func Test_NewName_InvalidName(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	invalidName := "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up "

	/// Act
	name, err := NewName(invalidName)

	/// Assert
	assert.Nil(name)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Name must not be greater than 50")
}
