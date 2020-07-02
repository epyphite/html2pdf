package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/epyphite/html2pdf/pkg/models"
	"github.com/epyphite/html2pdf/pkg/service"
	"github.com/epyphite/html2pdf/pkg/utils"
)

var rootCmd = &cobra.Command{
	Use:   "h2pdf",
	Short: "GO Html to pdf",
	Long:  ``,
	RunE:  runHTML2PDF,
}

//Execute will run the desire module command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var url string
var configFile string
var storeFolder string

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "download a single url")
	rootCmd.PersistentFlags().StringVar(&storeFolder, "storeFolder", "", "Specify a storage location")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Specify a configuration file")

}

func runHTML2PDF(cmd *cobra.Command, args []string) error {

	var config models.Config
	var srv service.HTML2PDF
	var err error

	if configFile == "" {
		config, err = utils.LoadConfigurationDefaults()

	} else {
		config, err = utils.LoadConfiguration(configFile)

	}
	srv.Setup(config)

	if url != "" {
		srv.GetURL(url)
	}
	if configFile != "" {
		urls, err := srv.GetURLFromFile()
		if err != nil {
			return err
		}
		for _, url := range urls {
			srv.GetURL(url)
		}
	}
	return err
}
