package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sbvalois/gojira/cmd"
	"github.com/sbvalois/gojira/setup"
	"github.com/spf13/viper"
)

const FILE_NAME string = "gojira.toml"

const PATH string = "."

func init() {
	viper.SetConfigName(FILE_NAME)
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/gojira")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("It looks like you've never ran the setup before")
			setup.RunBasicSetup(FILE_NAME, bufio.NewReader(os.Stdin))
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

}

func main() {
	app := cmd.CreateCliApp(FILE_NAME)

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Something went wrong, %v", err)
	}
}
