package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/sbvalois/gojira/helpers"
)

func getIssueUrl(issue string) string {
	space := helpers.GetJiraSpaceUri()
	return fmt.Sprintf("https://%s/browse/%s", space, issue)
}

func OpenBrowser(issue string) {
	var err error
	url := getIssueUrl(issue)

	fmt.Printf("Opening issue %s in browser...\n", issue)

	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		fmt.Println("Unsupported platform. Cannot open browser")
	}
	if err != nil {
		log.Fatalf("Something went wrong when opening browser: %v", err)
	}
}
