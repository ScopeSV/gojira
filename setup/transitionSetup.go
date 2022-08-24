package setup

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func RunTransitionSetup(filename string, r InputReader) {
	fmt.Println("In order to change issue status, you need to enter status transition ids.")
	fmt.Println("Unfortunately, there are no super easy way to access your transitions ids.")
	fmt.Println("The easiest way is to do a cURL request on a jira issue which is \"in progress\",")
	fmt.Println("as these issues usually have the most transition states.")
	fmt.Println("")
	fmt.Println("curl --request GET \\")
	fmt.Printf("  --url \"https://%s/rest/api/2/issue/{issue-key}/transitions\" \\\n", viper.GetString("jira.space"))
	fmt.Printf("  --user \"%s:%s\" \\\n", viper.GetString("user.email"),
		viper.GetString("jira.apiKey"),
	)
	fmt.Println("  --header \"Accept: application/json\" \\")
	fmt.Println("")
	fmt.Println("Gojira supports 3 transition states at the moment: \"open\", \"in progress\", \"review\" and \"done\"")
	fmt.Println("Please enter the transition id for each of these states:")

	createEntry("Open", "transitions.open", r)
	createEntry("In Progress", "transitions.inProgress", r)
	createEntry("Review", "transitions.review", r)
	createEntry("Done", "transitions.done", r)

	if err := viper.WriteConfigAs(filename); err != nil {
		log.Fatalf("Something went wrong when writing config file: %v", err)
	}

	SaveConfigGlobally(filename)

	fmt.Println("Transition Setup complete")
	fmt.Println("You may at any point run `gojira setup` to change your settings")
}
