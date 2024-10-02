package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Description of the project",
	Long: `
	Small description about the project develop in the technical test for Fast Track
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("**************")
		fmt.Println("This is my submition on the technical test for Fast Track.")
		fmt.Println("It has a REST API built using the framework Gin and uses Cobra for the CLI.")
		fmt.Println("")
		fmt.Println("If you have any question regarding this project, feel free to send me an email: simaotech@gmail.com")
		fmt.Println("**************")
	},
}

func init() {
	rootCmd.AddCommand(readmeCmd)
}
