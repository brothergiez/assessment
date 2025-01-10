package bookingsystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRoom(t *testing.T) {
	system := NewBookingSystem()

	err := system.AddRoom("room1", 5)
	assert.NoError(t, err)

	err = system.AddRoom("room1", 5)
	assert.Error(t, err)
	assert.Equal(t, "room already exists", err.Error())
}

func TestBookRoom(t *testing.T) {
	system := NewBookingSystem()

	err := system.AddRoom("room1", 5)
	assert.NoError(t, err)

	message, err := system.Book("room1", "2025-01-10 10:00", 2, 3)
	assert.NoError(t, err)
	assert.Equal(t, "Memesan room1 mulai jam 10:00 selama 2 jam", message)

	_, err = system.Book("room1", "2025-01-10 11:00", 1, 2)
	assert.Error(t, err)
	assert.Equal(t, "booking conflict detected", err.Error())

	message, err = system.Book("room1", "2025-01-10 12:00", 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, "Memesan room1 mulai jam 12:00 selama 1 jam", message)

	_, err = system.Book("room1", "2025-01-10 14:00", 1, 6)
	assert.Error(t, err)
	assert.Equal(t, "capacity exceeds room limit", err.Error())

	_, err = system.Book("room2", "2025-01-10 10:00", 2, 3)
	assert.Error(t, err)
	assert.Equal(t, "room does not exist", err.Error())

	_, err = system.Book("room1", "invalid time", 2, 3)
	assert.Error(t, err)
	assert.Equal(t, "invalid start time format", err.Error())
}

func TestBookingIntegration(t *testing.T) {
	system := NewBookingSystem()

	err := system.AddRoom("room1", 5)
	assert.NoError(t, err)
	err = system.AddRoom("room2", 3)
	assert.NoError(t, err)

	message, err := system.Book("room1", "2025-01-10 09:00", 2, 4)
	assert.NoError(t, err)
	assert.Equal(t, "Memesan room1 mulai jam 09:00 selama 2 jam", message)

	message, err = system.Book("room2", "2025-01-10 10:00", 1, 2)
	assert.NoError(t, err)
	assert.Equal(t, "Memesan room2 mulai jam 10:00 selama 1 jam", message)

	_, err = system.Book("room1", "2025-01-10 10:00", 1, 2)
	assert.Error(t, err)
	assert.Equal(t, "booking conflict detected", err.Error())
}
