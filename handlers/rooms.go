package handlers

import (
	"net/http"

	"github.com/byuoitav/common/db"
	"github.com/byuoitav/common/log"
	"github.com/labstack/echo"
)

//GetRoomsByBuilding returns all the rooms in a room
func GetRoomsByBuilding(context echo.Context) error {
	log.L.Debug("[room] Starting GetRoomsByBuilding...")

	buildingID := context.Param("building")

	rooms, err := db.GetDB().GetRoomsByBuilding(buildingID)

	if err != nil {
		log.L.Errorf("[room] An error occurred while getting all rooms in the room %s: %v", buildingID, err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	log.L.Debugf("[room] Successfully got all rooms in the room %s!", buildingID)
	return context.JSON(http.StatusOK, rooms)
}

// GetRoomByID returns all info about a room
func GetRoomByID(context echo.Context) error {
	log.L.Debug("[room] Starting GetRoomByID...")

	id := context.Param("room")

	room, err := db.GetDB().GetRoom(id)
	if err != nil {
		log.L.Errorf("[room] Failed to get the room %s : %v", id, err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	log.L.Debugf("[room] Successfully got the room %s!", room.ID)
	return context.JSON(http.StatusOK, room)
}

// GetAllRooms returns all rooms from the database.
func GetAllRooms(context echo.Context) error {
	log.L.Debug("[room] Starting GetAllRooms...")

	rooms, err := db.GetDB().GetAllRooms()
	if err != nil {
		log.L.Errorf("[room] Failed to get all rooms : %v", err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	log.L.Debug("[room] Successfully got all rooms!")
	return context.JSON(http.StatusOK, rooms)
}
