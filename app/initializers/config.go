package initializers

import "github.com/golly-go/golly"

func configPreboot() error {
	golly.SetName("skeleton")
	golly.SetVersion(0, 0, 1, "")

	// append our config initailizer to the first config slot always
	// you can manually do this in config.Initializers but for saftey i
	// chose todo this here to guarantee its always close to the first one
	golly.RegisterInitializerEx(true, configInitializer)
	return nil
}

// ConfigInitializer initializes various config options
func configInitializer(a golly.Application) error {
	a.Config.SetDefault("bind", ":9004")

	// Cors config
	a.Config.SetDefault("cors", map[string]interface{}{
		"origins": []string{
			"http://localhost:3000",
			"http://localhost:*",
			"http://127.0.0.1:*",
		},
		"headers": []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"Link",
		},
		"methods": []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
			"HEAD",
			"PATCH",
		},
	})
	return nil
}
