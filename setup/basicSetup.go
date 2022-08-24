package setup

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func RunBasicSetup(filename string, r InputReader) {
	fmt.Println("Setup")
	fmt.Println("=====")
	fmt.Println("We need to setup your Jira integration")
	fmt.Println("(Leave blank if you don't want to overwrite existing values)")
	fmt.Println("--------------")

	createEntry("Enter jira space (ex: foobar.atlassian.net)", "jira.space", r)
	createEntry("Enter email", "user.email", r)
	createEntry(
		"Enter API Key (You can generate it in your Jira account settings)",
		"jira.apiKey",
		r,
	)

	if err := viper.WriteConfigAs(filename); err != nil {
		log.Fatalf("Could not write config: %v", err)
	}

	SaveConfigGlobally(filename)

	if IsBasicSetupComplete() {
		fmt.Println("Basic Setup complete")
		fmt.Println("Do you want to setup your Jira transitions now? (y/n)")
		if ReadInput(r) == "y" {
			RunTransitionSetup(filename, r)
		}
	}

	os.Exit(0)
}
