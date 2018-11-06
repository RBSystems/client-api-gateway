package handlers

import (
	"net/http"

	"github.com/byuoitav/common/db"
	"github.com/byuoitav/common/log"
	"github.com/labstack/echo"
)

// GetBuildings returns a list of all the buildings in the database.
func GetBuildings(context echo.Context) error {
	log.L.Debug("[bldg] Starting GetBuildings...")

	buildings, err := db.GetDB().GetAllBuildings()
	if err != nil {
		log.L.Errorf("[bldg] Failed to get all buildings : %v", err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	log.L.Debug("[bldg] Successfully got all buildings!")
	return context.JSON(http.StatusOK, buildings)
}

// GetBuildingByID returns a specific building based on the given ID.
func GetBuildingByID(context echo.Context) error {
	log.L.Debug("[bldg] Starting GetBuildingByID...")

	id := context.Param("building")

	building, err := db.GetDB().GetBuilding(id)
	if err != nil {
		log.L.Errorf("[bldg] Failed to get the building %s : %v", id, err.Error())
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	log.L.Debugf("[bldg] Successfully got the building %s!", building.ID)
	return context.JSON(http.StatusOK, building)
}
