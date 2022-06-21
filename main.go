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

type IssueSprintType struct {
}

// TODO Kan lage ett flag som sorterer på date

var filename string = "config.toml"

func init() {
	viper.SetConfigName(filename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("It looks like you've never ran the setup before")
			setup.RunBasicSetup(filename, bufio.NewReader(os.Stdin))
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

}

func main() {
	app := cmd.CreateCliApp(filename)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Something went wrong, %v", err)
	}
}
