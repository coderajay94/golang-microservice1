package main

import (
	"log"
	"net/http"
	"time"

	"github.com/coderajay94/microservice1/base"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const(
	ServiceName = "useraccounts"
	ServiceVersion = "v1"
	ServiceOwner = "accounts"
	ServicePort = 8080
	PrometheusPort = 8082

	BasePath = "/accounts/"+ServiceVersion

	QueryLimit = 100
	ServerTimeout = 1 * time.Second

	DatabaseName = "user"
	CollectionName = "accounts"

)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("starting user accounts service...")
	
	err := godotenv.Load()
	if err != nil {
		logger.Panic("failed to load environment variables, aborting service initialization..")
	}

	var s base.Service
	{
		s = base.NewService(logger, QueryLimit)
		s = base.NewLoggingMiddleware(logger)(s)
	}
	endpoints := base.MakeServerEndpoints(s)

	h := base.MakeHttpHandler(endpoints, BasePath)



	// httpServer := http.Server{
	// 	Addr: ":"+strconv.Itoa(ServicePort),
	// 	Handler: h,
	// }

	log.Fatal(http.ListenAndServe(":8080", h))

	// go func(){
		
	// 	//logger.Info(fmt.Sprintf("starting the service on port %d", ServicePort))
	// 	//error := httpServer.ListenAndServe()
	// 	// if error != nil {
	// 	// 	logger.Error("Error starting the service"+ fmt.Sprintf("%d", error))
	// 	// }
	// }()

	logger.Info("stopped runing the service on port")
	
}

