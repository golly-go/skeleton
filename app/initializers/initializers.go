package initializers

import (
	"github.com/slimloans/golly"
	"github.com/slimloans/golly-skeleton/app/controllers"
)

// Preboots lists the preboots
var Preboots = []golly.PrebootFunc{
	func() error {
		golly.SetName("skeleton")
		golly.SetVersion(0, 0, 1, "")
		return nil
	},
}

// Initializers default app initializers - not sure if i like this yet
// id like eto keep the seperated for cleanliness
var Initializers = []golly.InitializerFunc{
	// orm.Initializer,
	func(a golly.Application) error { controllers.Routes(a); return nil },
}
