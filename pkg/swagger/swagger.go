package swagger

import "62tech.co/service/pkg/swagger/docs"

// @termsOfService http://swagger.io/terms/

func Init() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "62 TECHNOLOGY"
	docs.SwaggerInfo.Description = "code challenge"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
