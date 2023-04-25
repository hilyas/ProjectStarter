package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ProjectStarter",
	Short: "A CLI to bootstrap projects.",
	Long: `A CLI to start a new project and create the basic structure. 
	It will create a new directory with the name of the project and create the basic its structure.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
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

	if !isPatternValid(projectType, pattern) {
		return nil, fmt.Errorf("Pattern %s is not valid for project type %s", pattern, projectType)
	}

	var configPath string
	if isProjectTypeValid(projectType) {
		configPath = fmt.Sprintf("config/%s/%s.yml", projectType, pattern)
	} else {
		return nil, fmt.Errorf("Project type %s is not valid", projectType)
	}

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := viper.AllSettings()

	return config, nil
}

func isProjectTypeValid(projectType string) bool {
	validTypes := []string{"terraform", "ansible"}
	for _, t := range validTypes {
		if projectType == t {
			return true
		}
	}
	return false
}

// TODO: make sure that the pattern is valid for the project type
func isPatternValid(projectType string, pattern string) bool {
	if pattern == "" {
		return true
	}
	configPath := fmt.Sprintf("config/%s/%s.yml", projectType, pattern)
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// TODO: move to core/validators.go
func isCICDValid(cicd string) bool {
	validCICDs := []string{"github", "circle", "travis", "jenkins", "gitlab"}
	for _, c := range validCICDs {
		if cicd == c {
			return true
		}
	}
	return false
}

// TODO: move to core/creators.go
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

	err := os.MkdirAll(cicdPath, 0755)
	if err != nil {
		return err
	}
	return nil
}

func createNestedDirectories(basePath string, children []interface{}) error {
	for _, child := range children {
		childConfig := child.(map[string]interface{})
		path := filepath.Join(basePath, childConfig["name"].(string))

		// If the child has an extension, create a file; otherwise, create a directory.
		if filepath.Ext(path) != "" {
			_, err := os.Create(path)
			if err != nil {
				return err
			}
		} else {
			err := os.MkdirAll(path, 0755)
			if err != nil {
				return err
			}
		}

		if childDirs, ok := childConfig["children"]; ok {
			err := createNestedDirectories(path, childDirs.([]interface{}))
			if err != nil {
				return err
			}
		}
	}
	return nil
}


func createDirectoryStructure(projectName string, config map[string]interface{}, cicd string, tests bool) error {
	children := config["children"].([]interface{})
	if err := createNestedDirectories(projectName, children); err != nil {
		return err
	}

	if cicd != "" || isCICDValid(cicd) {
		err := createCICDFile(projectName, cicd)
		if err != nil {
			fmt.Println("CI/CD option not valid. Skipping creation.")
			return err
		}
	}

	if tests {
		testsPath := filepath.Join(projectName, "tests")
		err := os.MkdirAll(testsPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}