package main

import (
	"net/http"
	"os"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/central-event-system/hub/base"
	"github.com/byuoitav/central-event-system/messenger"
	"github.com/byuoitav/client-api-gateway/handlers"
	"github.com/byuoitav/client-api-gateway/socket"
	"github.com/byuoitav/common/log"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	mess, err := messenger.BuildMessenger(os.Getenv("HUB_ADDRESS"), base.Messenger, 1000)
	if err != nil {
		log.L.Fatalf("Failed to build event messenger : %s", err.String())
	}

	handlers.SetMessenger(mess)

	go handlers.WriteEventsToSocket()

	port := ":9100"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// websocket
	router.GET("/websocket", func(context echo.Context) error {
		socket.ServeWebsocket(context.Response().Writer, context.Request())
		return nil
	})

	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	router.PUT("/log-level/:level", log.SetLogLevel)
	router.GET("/log-level", log.GetLogLevel)

	secure.GET("/buildings", handlers.GetBuildings)
	secure.GET("/buildings/:building", handlers.GetBuildingByID)

	secure.GET("/buildings/:building/rooms", handlers.GetRoomsByBuilding)
	secure.GET("/rooms", handlers.GetAllRooms)
	secure.GET("/rooms/:room", handlers.GetRoomByID)

	secure.GET("/buildings/:building/rooms/:room", handlers.GetRoomState)
	secure.PUT("/buildings/:building/rooms/:room", handlers.SetRoomState)

	router.GET("/buildings/:building/rooms/:room/subscribe", handlers.SubscribeToRoom)
	router.GET("/buildings/:building/rooms/:room/unsubscribe", handlers.UnsubscribeFromRoom)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
