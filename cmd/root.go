/*
Copyright © 2023 Ilyas Hamdi <ilyas.hamdi@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ProjectStarter",
	Short: "A CLI to bootstrap projects.",
	Long: `A CLI to start a new project and create the basic structure. 
	It will create a new directory with the name of the project and create the basic its structure.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ProjectStarter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var cfgFile string
var homedir struct {
	Dir func() (string, error)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ProjectStarter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ProjectStarter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func readConfig(projectType string, pattern string) (map[string]interface{}, error) {
	if pattern == "" {
		pattern = "basic"
	}
	configPath := fmt.Sprintf("config/%s/%s.yml", projectType, pattern)

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := viper.AllSettings()

	return config, nil
}

// func isProjectTypeValid(projectType string) bool {
// 	validTypes := []string{"go", "python", "java", "javascript", "rust", "c", "c++", "c#", "php", "ruby", "scala", "swift", "kotlin", "dart", "elixir", "haskell", "clojure", "lua", "perl", "r", "fortran", "go", "groovy", "julia", "ocaml", "powershell", "racket", "tcl", "typescript", "assembly", "cobol", "erlang", "fsharp", "haxe", "nim", "pascal", "prolog", "smalltalk", "verilog", "zsh"}
// 	for _, t := range validTypes {
// 		if projectType == t {
// 			return true
// 		}
// 	}
// 	return false
// }

// func isPatternValid(projectType string, pattern string) bool {
// 	if pattern == "" {
// 		return true
// 	}
// 	configPath := fmt.Sprintf("config/%s/%s.yml", projectType, pattern)
// 	_, err := os.Stat(configPath)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return true
// }

// func isCICDValid(cicd string) bool {
// 	validCICDs := []string{"github", "circle", "travis", "jenkins", "gitlab"}
// 	for _, c := range validCICDs {
// 		if cicd == c {
// 			return true
// 		}
// 	}
// 	return false
// }

func createCICDFile(projectName string, cicd string) error {
	var cicdPath string
	switch cicd {
	case "github":
		cicdPath = filepath.Join(projectName, ".github", "workflows", "main.yml")
	case "circle":
		cicdPath = filepath.Join(projectName, ".circleci", "config.yml")
	case "travis":
		cicdPath = filepath.Join(projectName, ".travis.yml")
	case "jenkins":
		cicdPath = filepath.Join(projectName, "Jenkinsfile")
	case "gitlab":
		cicdPath = filepath.Join(projectName, ".gitlab-ci.yml")
	}

	_, err := os.Create(cicdPath)
	if err != nil {
		return err
	}
	return nil
}

func createTestFile(projectName string, projectType string) error {
	var testPath string
	switch projectType {
	case "terraform":
		testPath = filepath.Join(projectName, "tests", "main.tf")
	case "ansibe":
		testPath = filepath.Join(projectName, "tests", "main.yml")
	default:
		testPath = filepath.Join(projectName, "tests", "test_file.txt")
	}

	_, err := os.Create(testPath)
	if err != nil {
		return err
	}

	return nil
}


func createNestedDirectories(basePath string, dirs []interface{}) error {
	for _, dir := range dirs {
		dirConfig := dir.(map[string]interface{})
		path := filepath.Join(basePath, dirConfig["name"].(string))
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}

		if children, ok := dirConfig["children"]; ok {
			childDirs := children.([]interface{})
			if err := createNestedDirectories(path, childDirs); err != nil {
				return err
			}
		}
	}
	return nil
}

func createDirectoryStructure(projectName string, config map[string]interface{}, cicd string, tests bool) error {
	directories := config["directories"].([]interface{})
	// Create directories from the config file
	if err := createNestedDirectories(projectName, directories); err != nil {
		return err
	}

	files := config["files"].([]interface{})
	for _, file := range files {
		fileConfig := file.(map[string]interface{})
		filePath := filepath.Join(projectName, fileConfig["name"].(string))
		_, err := os.Create(filePath)
		if err != nil {
			return err
		}
	}

	if cicd != "" {
		createCICDFile(projectName, cicd)
	// var cicdPath string
	// 	switch cicd {
	// 	case "github":
	// 		cicdPath = filepath.Join(projectName, ".github", "workflows", "main.yml")
	// 	case "circle":
	// 		cicdPath = filepath.Join(projectName, ".circleci", "config.yml")
	// 	case "travis":
	// 		cicdPath = filepath.Join(projectName, ".travis.yml")
	// 	case "jenkins":
	// 		cicdPath = filepath.Join(projectName, "Jenkinsfile")
	// 	case "gitlab":
	// 		cicdPath = filepath.Join(projectName, ".gitlab-ci.yml")
	// 	default:
	// 		fmt.Println("Invalid CI/CD option. Skipping CI/CD directory creation.")
	// 	} // TODO: Some systems use a file instead of a directory: Jenkins, Gitlab, Travis

	// 	if cicdPath != "" {
	// 		err := os.MkdirAll(cicdPath, 0755)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	}

	if tests {
		createTestFile(projectName, config["type"].(string))
		// testsPath := filepath.Join(projectName, "tests")
		// err := os.MkdirAll(testsPath, 0755)
		// if err != nil {
		// 	return err
		// }
	}

	return nil
}

