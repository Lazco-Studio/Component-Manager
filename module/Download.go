package module

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/gookit/color"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func Download(url string, destination string) string {
	client := grab.NewClient()
	req, _ := grab.NewRequest(destination, url)

	resp := client.Do(req)
	filePath := resp.Filename
	fileName := filepath.Base(filePath)

	fmt.Printf("Package source: ")
	color.Cyanln(req.URL())
	fmt.Printf("Saving as:      ")
	color.Cyanln(fileName)
	ProgressBar(resp)

	if err := resp.Err(); err != nil {
		log.Fatal("Download failed")
	}

	return filePath
}

func ProgressBar(resp *grab.Response) {
	status := ""
	if resp.HTTPResponse.Status == "200 OK" {
		status = "Downloading"
	} else if resp.HTTPResponse.Status == "206 Partial Content" {
		status = "Resuming Download"
	}

	bar := progressbar.NewOptions(
		int(resp.Size()),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()), //you should install "github.com/k0kubun/go-ansi"
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSpinnerType(1),
		progressbar.OptionShowBytes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetDescription("[magenta]"+status+"[reset]"),
		progressbar.OptionOnCompletion(func() {
			color.Greenln("\nDownload complete!")
		}),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]█[reset]",
			SaucerHead:    "[green]▓▒░[reset]",
			SaucerPadding: " ",
			BarStart:      "▐",
			BarEnd:        "▌",
		}))

	t := time.NewTicker(5 * time.Millisecond)
	defer t.Stop()
Loop:
	for {
		select {
		case <-t.C:
			bar.Set(int(resp.BytesComplete()))

		case <-resp.Done:
			bar.Set(int(resp.Size()))
			fmt.Println()
			break Loop
		}
	}
}
