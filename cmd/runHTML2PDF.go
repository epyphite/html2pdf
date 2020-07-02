package cmd
import (
	"context"
	"fmt"
	"os/exec"
	"path"
	"time"
	"github.com/spf13/cobra"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)


var rootCmd = &cobra.Command{
	Use:   "go-h2pdf",
	Short: "go Html to pdf",
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
var config string
var storeFolder string

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "download a single url")
	rootCmd.PersistentFlags().StringVar(&storeFolder, "storeFolder", "", "Specify a storage location")

}


func getURLFromFile(file string) ([]string, error) {
	var returnURLs []string
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return returnURLs, err
	}
	// Get value from cell by given worksheet name and axis.
	rows, err := f.GetRows("FDA-Warning-Letters-2-Parse-Json")
	for _, row := range rows {
		for i, colCell := range row {
			if i == 3 {
				returnURLs = append(returnURLs, colCell)
			}
		}
	}

	return returnURLs, err
}

func getURL(url string) {
	fmt.Println("Getting URL ", url)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	fname := storeFolder +"/" + path.Base(url) + ".pdf"
	cmd := exec.CommandContext(
		ctx,
		"latest/chrome",
		//	"--no-sandbox",
		"--disable-gpu",
		//	"--virtual-time-budget=2000",
		//	"--timeout=6000",
		"--headless",
		fmt.Sprintf("--print-to-pdf=%s", fname),
		fmt.Sprintf(
			"%s", url),
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func setup(storeFolder string) {

	if storeFolder == "" {
		storeFolder = "./pdfStore/"
	}
	_, err := os.Stat(storeFolder)
	if err != nil {
		os.IsNotExist(err)
	}
	err = os.Mkdir(storeFolder, 0770)
	if err != nil {
		fmt.Println((err))
	}

}

func runHTML2PDF(cmd *cobra.Command, args []string) error {

	setup(storeFolder)
	var err error
	if url != "" {
		getURL(url)
	}
	if config != "" {
		urls, err := getURLFromFile(config)
		if err != nil {
			return err
		}
		for _, url := range urls {
			getURL(url)
		}
	}
	return err
}