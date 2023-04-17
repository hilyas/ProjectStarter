package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project structure",
	Long:  `Create a new project structure based on the specified project type, structure pattern, and optional flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectType, _ := cmd.Flags().GetString("type")
		projectName, _ := cmd.Flags().GetString("name")
		pattern, _ := cmd.Flags().GetString("pattern")
		boilerplate, _ := cmd.Flags().GetBool("boilerplate")
		cicd, _ := cmd.Flags().GetString("cicd")
		tests, _ := cmd.Flags().GetBool("tests")
	
		// Read the config file for the type and pattern
		config, err := readConfig(projectType, pattern)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		// Create the project directory
		err = os.Mkdir(projectName, 0755)
		if err != nil {
			fmt.Println("Error creating project directory:", err)
			return
		}

		// Create the directory structure based on the config file
		err = createDirectoryStructure(projectName, config, cicd, tests)
		if err != nil {
			fmt.Println("Error creating directory structure:", err)
			return
		}

		// TODO: Optionally add a CI/CD directory if the flag is set, separate it from createDirectoryStructure
	
		// TODO: Optionally add boilerplate code if the flag is set, separate it from createDirectoryStructure
	
		// TODO: Add a file flag to take an external config file to create the project structure

		// use the variables to avoid compiler errors
		fmt.Println(projectType, pattern, boilerplate, cicd, tests)
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
	createCmd.Flags().StringP("name", "n", "", "Project name (optional)")
	createCmd.Flags().StringP("pattern", "p", "", "Structure pattern (optional)")
	createCmd.Flags().BoolP("boilerplate", "b", false, "Include boilerplate code (optional)")
	createCmd.Flags().StringP("cicd", "c", "", "Add a CI/CD directory (optional: github, circle, travis, jenkins, gitlab)")
	createCmd.Flags().BoolP("tests", "s", false, "Add a test directory (optional)")
	createCmd.MarkFlagRequired("type")
}


