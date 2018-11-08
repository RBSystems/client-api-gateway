package main

import (
	"net/http"
	"os"

	"github.com/byuoitav/central-event-system/hub/base"
	"github.com/byuoitav/central-event-system/messenger"
	"github.com/byuoitav/client-api-gateway/handlers"
	"github.com/byuoitav/client-api-gateway/socket"
	"github.com/byuoitav/common"
	"github.com/byuoitav/common/log"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
)

func main() {
	mess, err := messenger.BuildMessenger(os.Getenv("HUB_ADDRESS"), base.Messenger, 1000)
	if err != nil {
		log.L.Fatalf("Failed to build event messenger : %s", err.String())
	}

	handlers.SetMessenger(mess)

	go handlers.WriteEventsToSocket()

	port := ":9100"
	router := common.NewRouter()

	// websocket
	router.GET("/websocket", func(context echo.Context) error {
		socket.ServeWebsocket(context.Response().Writer, context.Request())
		return nil
	})

	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	router.PUT("/log-level/:level", log.SetLogLevel)
	router.GET("/log-level", log.GetLogLevel)

	router.GET("/buildings", handlers.GetBuildings)
	router.GET("/buildings/:building", handlers.GetBuildingByID)

	router.GET("/buildings/:building/rooms", handlers.GetRoomsByBuilding)
	router.GET("/rooms", handlers.GetAllRooms)
	router.GET("/rooms/:room", handlers.GetRoomByID)

	router.GET("/buildings/:building/rooms/:room", handlers.GetRoomState)
	router.PUT("/buildings/:building/rooms/:room", handlers.SetRoomState)

	router.GET("/buildings/:building/rooms/:room/subscribe", handlers.SubscribeToRoom)
	router.GET("/buildings/:building/rooms/:room/unsubscribe", handlers.UnsubscribeFromRoom)

	router.Static("/", "lookout-dist")

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
