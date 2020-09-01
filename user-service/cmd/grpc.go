package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"

	handlers "user-service/internal/handlers/grpc"
	"user-service/internal/repositories"
	"user-service/internal/services"
	"user-service/pkg/pb"
)

func ServeGRPC() *cobra.Command {
	return &cobra.Command{
		Use: "grpc:serve",
		Run: func(cmd *cobra.Command, args []string) {
			userRepo := repositories.NewJSONPlaceholderUserRepository()
			userSvc := services.NewUserService(userRepo)
			userHandler := handlers.NewUserHandler(userSvc)

			server := grpc.NewServer()
			pb.RegisterUserServiceServer(server, userHandler)

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
