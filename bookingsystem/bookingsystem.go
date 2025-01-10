package bookingsystem

import (
	"errors"
	"fmt"
	"time"
)

type BookingSystem struct {
	rooms map[string]*Room
}

type Room struct {
	ID       string
	Capacity int
	Bookings []Booking
}

type Booking struct {
	StartTime time.Time
	EndTime   time.Time
	Capacity  int
}

func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		rooms: make(map[string]*Room),
	}
}

func (bs *BookingSystem) AddRoom(id string, capacity int) error {
	if _, exists := bs.rooms[id]; exists {
		return errors.New("room already exists")
	}
	bs.rooms[id] = &Room{
		ID:       id,
		Capacity: capacity,
		Bookings: []Booking{},
	}
	return nil
}

func (bs *BookingSystem) Book(roomID, startTimeStr string, durationHours, capacity int) (string, error) {
	room, exists := bs.rooms[roomID]
	if !exists {
		return "", fmt.Errorf("room does not exist")
	}
	if capacity > room.Capacity {
		return "", fmt.Errorf("capacity exceeds room limit")
	}

	startTime, err := time.Parse("2006-01-02 15:04", startTimeStr)
	if err != nil {
		return "", fmt.Errorf("invalid start time format")
	}
	endTime := startTime.Add(time.Duration(durationHours) * time.Hour)

	for _, booking := range room.Bookings {
		if startTime.Before(booking.EndTime) && endTime.After(booking.StartTime) {
			return "", fmt.Errorf("booking conflict detected")
		}
	}

	room.Bookings = append(room.Bookings, Booking{
		StartTime: startTime,
		EndTime:   endTime,
		Capacity:  capacity,
	})

	return fmt.Sprintf("Memesan %s mulai jam %s selama %d jam", roomID, startTime.Format("15:04"), durationHours), nil
}
