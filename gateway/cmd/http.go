package cmd

import (
	"log"
	"net/http"
	"os"

	handlers "gateway/internal/handlers/http"
	"gateway/internal/repositories"
	"gateway/internal/services"
	"gateway/pkg/pb"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func ServeHTTP() *cobra.Command {
	return &cobra.Command{
		Use: "http:serve",
		Run: func(cmd *cobra.Command, args []string) {
			userGrpcConn, err := grpc.Dial(
				os.Getenv("MS_USER_SERVICE_HOST")+":"+os.Getenv("MS_USER_SERVICE_GRPC_PORT"),
				grpc.WithInsecure(),
				grpc.WithBlock(),
			)
			if err != nil {
				log.Fatalln(err)
			}
			userGrpc := pb.NewUserServiceClient(userGrpcConn)

			newsGrpcConn, err := grpc.Dial(
				os.Getenv("MS_NEWS_SERVICE_HOST")+":"+os.Getenv("MS_NEWS_SERVICE_GRPC_PORT"),
				grpc.WithInsecure(),
				grpc.WithBlock(),
			)
			if err != nil {
				log.Fatalln(err)
			}
			newsGrpc := pb.NewNewsServiceClient(newsGrpcConn)

			newsRepo := repositories.NewMSNewsRepository(newsGrpc)
			userRepo := repositories.NewMSUserRepository(userGrpc)
			newsSvc := services.NewNewsService(userRepo, newsRepo)

			router := prepareHTTPRouter(newsSvc)
			if err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), router); err != nil {
				log.Fatalln(err)
			}
		},
	}
}

func prepareHTTPRouter(newsSvc services.NewsService) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello, world!"))
	})

	router.Mount("/news", handlers.NewNewsHandler(newsSvc).GetRoutes())

	return router
}
