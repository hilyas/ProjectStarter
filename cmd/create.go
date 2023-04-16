package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project structure",
	Long:  `Create a new project structure based on the specified project type, structure pattern, and optional flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectType, _ := cmd.Flags().GetString("type")
		pattern, _ := cmd.Flags().GetString("pattern")
		boilerplate, _ := cmd.Flags().GetBool("boilerplate")
		cicd, _ := cmd.Flags().GetBool("cicd")
	
		// TODO: Read the config file for the type and pattern
	
		// TODO: Create the directory structure based on the config file
	
		// TODO: Optionally add a CI/CD directory if the flag is set
	
		// TODO: Optionally add boilerplate code if the flag is set
	
		// use the variables to avoid compiler errors
		fmt.Println(projectType, pattern, boilerplate, cicd)
		fmt.Println("Project created successfully")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	createCmd.Flags().StringP("type", "t", "", "Project type (required)")
	createCmd.Flags().StringP("pattern", "p", "", "Structure pattern (optional)")
	createCmd.Flags().BoolP("boilerplate", "b", false, "Include boilerplate code (optional)")
	createCmd.Flags().BoolP("cicd", "c", false, "Include CI/CD configuration (optional)")
	createCmd.MarkFlagRequired("type")
}


