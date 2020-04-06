package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: zoom-handler ZOOM_URL")
		os.Exit(0)
	}

	joinURL := os.Args[1]

	// Check if link is from a google calendar link
	if strings.Contains(joinURL, "google.com/url?q=") {
		joinURL = extractZoomURL(joinURL)
	}

	parsedURL, err := url.Parse(joinURL)
	if err != nil {
		log.Fatal(err)
	}

	queryParams, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}

	meetingID := path.Base(parsedURL.Path)
	password := queryParams.Get("pwd")
	openURL := fmt.Sprintf("zoommtg://zoom.us/join?action=join&confno=%s&pwd=%s", meetingID, password)

	exec.Command("xdg-open", openURL).Run()
}

func extractZoomURL(URL string) string {
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}

	queryParams, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Fatal(err)
	}

	embeddedURL := queryParams.Get("q")
	if len(embeddedURL) == 0 {
		err = fmt.Errorf("Invalid zoom link in %s", URL)
		log.Fatal(err)
	}

	return embeddedURL
}
