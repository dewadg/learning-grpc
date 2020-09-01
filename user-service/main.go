package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"user-service/cmd"
)

var rootCMD = &cobra.Command{
	Use:   "api",
	Short: "Entrypoint for running this service",
}

func init() {
	godotenv.Load()

	rootCMD.AddCommand(cmd.ServeGRPC())
}

func main() {
	rootCMD.Execute()
}
