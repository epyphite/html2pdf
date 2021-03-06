package service

import (
	"context"
	"os/exec"
	"path"

	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/epyphite/html2pdf/pkg/models"
)

//HTML2PDF Main service structure
type HTML2PDF struct {
	Config models.Config
}

//Setup Will setup basic directories
func (H2 *HTML2PDF) Setup(config models.Config) {
	H2.Config = config
	_, err := os.Stat(H2.Config.DestinationFolder)
	if err != nil {
		os.IsNotExist(err)
	}
	err = os.Mkdir(H2.Config.DestinationFolder, 0770)
	if err != nil {
		log.Println(err)
	}

}

//GetURLFromFile Will get url from the specified format file
func (H2 *HTML2PDF) GetURLFromFile() ([]string, error) {
	var returnURLs []string
	var err error

	if H2.Config.SourceType == "XLS" {
		f, err := excelize.OpenFile(H2.Config.SourceFolder + H2.Config.SourceName)
		if err != nil {
			fmt.Println(err)
			return returnURLs, err
		}
		// Get value from cell by given worksheet name and axis.
		rows, err := f.GetRows(H2.Config.TabName)
		for _, row := range rows {
			for i, colCell := range row {
				col, err := strconv.Atoi(H2.Config.ColumnNumber)
				if err != nil {
					log.Println(err)
				}
				if i == col {
					returnURLs = append(returnURLs, colCell)
				}
			}
		}
	}

	return returnURLs, err
}

//GetURL gets the desire url and converts it into a PDF saving it to the destination folder
func (H2 *HTML2PDF) GetURL(url string) {
	fmt.Println("Getting URL ", url)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fname := H2.Config.DestinationFolder + "/" + path.Base(url) + ".pdf"
	cmd := exec.CommandContext(
		ctx,
		"latest/chrome",
		"--disable-gpu",
		"--headless",
		fmt.Sprintf("--print-to-pdf=%s", fname),
		fmt.Sprintf(
			"%s", url),
	)

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
