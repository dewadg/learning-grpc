package cmd

import (
	"log"
	"net"
	"news-service/pkg/pb"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	handlers "news-service/internal/handlers/grpc"
	"news-service/internal/repositories"
	"news-service/internal/services"
)

func ServeGRPC() *cobra.Command {
	return &cobra.Command{
		Use: "grpc:serve",
		Run: func(cmd *cobra.Command, args []string) {
			newsRepo := repositories.NewJSONPlaceholderNewsRepository()
			newsSvc := services.NewNewsService(newsRepo)
			newsHandler := handlers.NewNewsHandler(newsSvc)

			server := grpc.NewServer()
			pb.RegisterNewsServiceServer(server, newsHandler)

			listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
			if err != nil {
				log.Fatalln(err)
			}

			if err := server.Serve(listener); err != nil {
				log.Fatalln(err)
			}
		},
	}
}
