package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

var (
	// Configuration holds configuration read from config.yaml.
	Configuration struct {
		Data map[string]interface{} `yaml:"data"`
	}
)

func init() {
	err := Configure()
	if err != nil {
		log.Println("config.Configure():", err.Error())
		os.Exit(3)
	}
}

// Configure parses the command line and loads system configuration from the
// configuration file.
func Configure() error {
	// Parse the command line for options.
	showHelp, configurationFile := parseCommandLine()

	// Print usage?
	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	// Load config.yaml.
	err := loadConfiguration(configurationFile)
	if err != nil {
		return fmt.Errorf("config.loadConfiguration(): %s", err)
	}

	return nil
}

func loadConfiguration(configurationFilePath string) error {
	// Check if a path to configuration file was provided.
	if configurationFilePath == "" {
		// Configuration file is mandatory.
		return fmt.Errorf("configuration file is mandatory and missing")
	}

	// Parse the configuration file in order to get the system configuration.
	fileContents, err := ioutil.ReadFile(configurationFilePath)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile(%s): %s", configurationFilePath, err)
	}
	err = yaml.Unmarshal(fileContents, &Configuration)
	if err != nil {
		return fmt.Errorf("yaml.Unmarshal(%s): %s", configurationFilePath, err)
	}

	return nil
}

func parseCommandLine() (printUsage bool, configurationFile string) {
	flag.BoolVar(&printUsage, "help", false, "print usage")
	flag.StringVar(&configurationFile, "config", "", "specify config.yaml file")
	flag.Parse()
	return printUsage, configurationFile
}
