package setup

import (
	"log"
	"os"
)

func SaveConfigGlobally(filename string) {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	os.MkdirAll(dirname+"/.config/gojira", 0700)
	if err := os.WriteFile(dirname+"/.config/gojira/"+filename, file, 0644); err != nil {
		log.Fatalf("Could not write config: %v", err)
	}
}
