package initializers

import "github.com/slimloans/golly"

// ConfigInitializer initializes various config options
func configInitializer(a golly.Application) error {
	a.Config.SetDefault("bind", ":9002")

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
