package main

import (
	"testing"

	core "github.com/zokypesch/ceria/core"

	mockMenu "github.com/zokypesch/example-ceria/command/mocks"

	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	mockLib "github.com/stretchr/testify/mock"
)

func TestConnection(t *testing.T) {
	assert := assert.New(t)

	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true
	// GORM
	_, err := gorm.Open(mocket.DriverName, "connection_string") // Can be any connection string

	assert.Equal(err, nil)
}

func TestMenu(t *testing.T) {
	conn := core.GetTestConnection()
	mock := new(mockMenu.MenuInterface)
	newMenu := InitialMenu(conn)

	mock.On("Menu", "db:migrate", conn).Return()
	mock.On("Contains", mockLib.AnythingOfType("string")).Return(true)

	mock.Menu("db:migrate", conn)

	exp := mock.Contains("db:migrate")
	act := newMenu.Contains("db:migrate")

	assert.Equal(t, exp, act)
	mock.AssertNumberOfCalls(t, "Menu", 1)

	mockConditionFalse := new(mockMenu.MenuInterface)
	mockConditionFalse.On("Contains", mockLib.AnythingOfType("string")).Return(false)

	expFalse := mockConditionFalse.Contains("fire")
	actFalse := newMenu.Contains("fire")

	// call the menu
	newMenu.Menu("db:migrate")

	// call the failed menu
	newMenu.Menu("db:rollback")

	// call main and hope run as well
	// main()

	assert.Equal(t, expFalse, actFalse)
	mock.AssertNumberOfCalls(t, "Contains", 1)
	mockConditionFalse.AssertNumberOfCalls(t, "Contains", 1)
	mock.AssertExpectations(t)
	mockConditionFalse.AssertExpectations(t)

}
