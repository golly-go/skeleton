package initializers

import (
	"github.com/slimloans/golly"
	"github.com/slimloans/golly-skeleton/app/controllers"
)

// RouterInitializer initializes our built in routes not sure
// if i like this as its just 1 line method
func routerInitializer(a golly.Application) error {
	controllers.Routes(a)
	return nil
}
