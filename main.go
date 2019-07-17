package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: zoom-handler ZOOM_URL")
		os.Exit(0)
	}

	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	b := path.Base(u.Path)
	s := fmt.Sprintf("zoommtg://zoom.us/join?action=join&confno=%s", b)

	exec.Command("xdg-open", s).Run()
}
