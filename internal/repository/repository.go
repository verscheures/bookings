package repository

import (
	"time"

	"github.com/verscheures/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilitByDatesByRoomID(start, end time.Time, roomId int) (bool, error)
	SearchAvailabilitForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomById(id int) (models.Room, error)
}
