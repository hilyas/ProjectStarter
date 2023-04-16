package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing project structure",
	Long: `Delete an existing project structure based on the specified project type, structure pattern, and optional flags.
Please be cautious when using this command, as it will delete files and directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		fmt.Println("WARNING: This will delete files and directories permanently. Please ensure you have backups if necessary.")
		fmt.Print("Are you sure you want to continue? (yes/no): ")
		var response string
		fmt.Scanln(&response)
		if response != "yes" {
			fmt.Println("Aborted.")
			return
		}

		// TODO: Read the config file for the type and pattern

		// TODO: Remove the directory structure based on the config file

		// use the variables to avoid compiler errors
		fmt.Println(name)
		fmt.Println("Project deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	deleteCmd.Flags().StringP("name", "n", "", "Project name (required)")
	deleteCmd.MarkFlagRequired("name")
}
