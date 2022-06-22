package setup

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type InputReader interface {
	ReadString(delim byte) (string, error)
}

func ReadInput(r InputReader) string {
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func IsBasicSetupComplete() bool {
	return viper.GetString("jira.space") != "" &&
		viper.GetString("user.email") != "" &&
		viper.GetString("jira.apiKey") != ""
}

func addToConfig(key string, value string) {
	if value != "" {
		viper.Set(key, value)
	}
}

func createEntry(text string, key string, r InputReader) {
	fmt.Printf("%s:\n", text)
	value := ReadInput(r)
	addToConfig(key, value)
}
