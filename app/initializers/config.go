package initializers

import "github.com/slimloans/golly"

func configPreboot() error {
	golly.SetName("skeleton")
	golly.SetVersion(0, 0, 1, "")

	golly.RegisterInitializerEx(true, configInitializer)
	return nil
}

// ConfigInitializer initializes various config options
func configInitializer(a golly.Application) error {
	a.Config.Set("bind", ":9004")

	// Cors config
	a.Config.Set("cors", map[string]interface{}{
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
