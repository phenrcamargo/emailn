package entity

import (
	"emailn/internal/campaign/application/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Contact_with_success(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	input := contract.NewContact{
		FirstName: "Test",
		LastName:  "Test",
		Phone:     "999999999",
		Email:     "teste@email.com",
	}
	/// Act
	contact, err := NewContact(input)

	/// Assert
	assert.Nil(err)
	assert.NotNil(contact)
	assert.NotNil(contact.ID)
	assert.EqualValues(*contact.FirstName, input.FirstName)
	assert.EqualValues(*contact.LastName, input.LastName)
	assert.EqualValues(*contact.Phone, input.Phone)
	assert.EqualValues(*contact.Email, input.Email)
}

func Test_Create_Contact_with_First_Name_Empty_error(t *testing.T) {
	assert := assert.New(t)
	input := contract.NewContact{
		FirstName: "",
		LastName:  "Test",
		Phone:     "999999999",
		Email:     "teste@email.com",
	}
	/// Act
	contact, err := NewContact(input)

	/// Assert
	assert.Nil(contact)
	assert.NotNil(err)
	assert.Contains(err.Error(), "First Name must not be empty")
}

func Test_Create_Contact_with_Last_Name_Empty_error(t *testing.T) {
	assert := assert.New(t)
	input := contract.NewContact{
		FirstName: "Test",
		LastName:  "",
		Phone:     "999999999",
		Email:     "teste@email.com",
	}
	/// Act
	contact, err := NewContact(input)

	/// Assert
	assert.Nil(contact)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Last Name must not be empty")
}

func Test_Create_Contact_with_Phone_Number_Empty_error(t *testing.T) {
	assert := assert.New(t)
	input := contract.NewContact{
		FirstName: "Test",
		LastName:  "Test",
		Phone:     "",
		Email:     "teste@email.com",
	}
	/// Act
	contact, err := NewContact(input)

	/// Assert
	assert.Nil(contact)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Phone Number must not be empty")
}

func Test_Create_Contact_with_Email_Empty_error(t *testing.T) {
	assert := assert.New(t)
	input := contract.NewContact{
		FirstName: "Test",
		LastName:  "Test",
		Phone:     "999999999",
		Email:     "",
	}
	/// Act
	contact, err := NewContact(input)

	/// Assert
	assert.Nil(contact)
	assert.NotNil(err)
	assert.Contains(err.Error(), "E-Mail Address must not be empty")
}
