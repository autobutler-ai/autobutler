package main

import (
	"fmt"
	"os"

	"dotai-go-backend/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "dotai"}
	rootCmd.AddCommand(cmd.Serve())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
