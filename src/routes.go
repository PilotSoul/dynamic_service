package src

import (
	controllers "PilotSoul/dynamic_service/src/interface/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// segments
	app.Post("/create_segment", controllers.CreateSegment)
	app.Post("/delete_segment", controllers.DeleteSegment)

	// users
	app.Post("/create_user", controllers.CreateUser)
	app.Post("/add_user_to_segment", controllers.AddSegments)
	app.Post("/show_segments/:user<int>", controllers.ShowUserSegments)
	app.Post("/delete_user_from_segment", controllers.DeleteSegments)
}
