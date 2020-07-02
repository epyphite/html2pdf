package utils

import (
	"encoding/json"
	"os"

	"github.com/epyphite/html2pdf/pkg/constants"
	"github.com/epyphite/html2pdf/pkg/models"
)

//LoadConfiguration returns the read Configuration and error while reading.
func LoadConfiguration(file string) (models.Config, error) {
	var config models.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}

func LoadConfigurationDefaults() (models.Config, error) {
	var config models.Config
	var err error
	config.ColumnNumber = constants.ColumnNumber
	config.DestinationFolder = constants.DestinationFolder
	config.SourceFolder = constants.SourceFolder
	config.SourceName = constants.SourceName
	config.SourceType = constants.SourceType
	config.TabName = constants.TabName

	return config, err
}
