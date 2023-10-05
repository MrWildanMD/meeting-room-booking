package routers

import (
	"github.com/MrWildanMD/room-booking/controllers"
	"github.com/MrWildanMD/room-booking/database"
	"github.com/MrWildanMD/room-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	booking := &controllers.BookingController{
		DB: database.DB,
	}
	users := &controllers.UsersController{
		DB: database.DB,
	}
	notif := &controllers.NotificationController{
		DB: database.DB,
	}
	room := &controllers.RoomController{
		DB: database.DB,
	}
	office := &controllers.OfficeController{
		DB: database.DB,
	}

	// Since booking doesnt require to log in for guest so it not necessary to use deserialize user middleware here
	bookingRoute := rg.Group("booking")
	{
		bookingRoute.POST("/", booking.AddBooking)
		bookingRoute.GET("/", booking.GetBooking)
		bookingRoute.GET("/:id", booking.GetBookingByID)
		bookingRoute.PUT("/:id", booking.UpdateBooking)
		bookingRoute.DELETE("/:id", booking.DeleteBooking)
	}

	notifRoute := rg.Group("notifications")
	{
		notifRoute.POST("/", notif.AddNotification)
		notifRoute.GET("/", notif.GetNotification)
		notifRoute.GET("/:id", notif.GetNotificationByID)
		notifRoute.PUT("/:id", notif.UpdateNotification)
		notifRoute.DELETE("/:id", notif.DeleteNotification)
	}

	officeRoute := rg.Group("office")
	{
		officeRoute.POST("/", office.AddOffice)
		officeRoute.GET("/", office.GetOffice)
		officeRoute.GET("/:id", office.GetOfficeByID)
		officeRoute.PUT("/:id", office.UpdateOffice)
		officeRoute.DELETE("/:id", office.DeleteOffice)
	}

	// Here we use our deserialize user middleware because admin can perform CRUD operations
	roomRoute := rg.Group("room")
	{
		roomRoute.POST("/", middlewares.DeserializeUsers(), room.AddRoom)
		roomRoute.GET("/", room.GetRoom)
		roomRoute.GET("/:id", room.GetRoomByID)
		roomRoute.PUT("/:id", middlewares.DeserializeUsers(), room.UpdateRoom)
		roomRoute.DELETE("/:id", middlewares.DeserializeUsers(), room.DeleteRoom)
	}

	usersRoute := rg.Group("users")
	{
		usersRoute.POST("/login", users.LoginUsers)
		usersRoute.POST("/register", users.RegisterUsers)
		usersRoute.POST("/logout", users.LogoutUsers)
		usersRoute.GET("/", middlewares.DeserializeUsers(), users.GetUsers)
		usersRoute.GET("/me", middlewares.DeserializeUsers(), users.GetMe)
		usersRoute.PUT("/:id", users.UpdateUsers)
		usersRoute.DELETE("/:id", middlewares.DeserializeUsers(), users.DeleteUsers)
	}
}
