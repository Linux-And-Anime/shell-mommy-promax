package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Config struct {
	Name    string   `json:"name"`
	Dialogs []string `json:"dialogs"`
	Pronoun string   `json:"pronoun"`
}

const Mommy = "mommy"

func main() {
	var config_dir string
	if path := os.Getenv("XDG_CONFIG_HOME"); path != "" {
		config_dir = fmt.Sprintf("%v/%v", path, Mommy)
	} else {
		if home_path := os.Getenv("HOME"); home_path != "" {
			config_dir = fmt.Sprintf("%v/.config/%v", home_path, Mommy)
		} else {
			log.Fatal("could not find config directory!")
		}
	}

	if err := os.Mkdir(config_dir, 0750); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	var config_path string = fmt.Sprintf("%v/%v", config_dir, "config.json")

	config_blob, err := os.ReadFile(config_path)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(config_path)
			if err := os.WriteFile(config_path, []byte("{}\n"), 0o660); err != nil {
				log.Fatal(err)
			}
			config_blob = []byte("{}")
		}
	}

	var config Config
	if err := json.Unmarshal(config_blob, &config); err != nil {
		log.Fatal("invalid config file: ", err)
	}

	if config.Name == "" {
		config.Name = Mommy
	}

	if len(config.Dialogs) == 0 {
		// TODO: change this things to some more reasonable things
		config.Dialogs = []string{
			fmt.Sprintf("%v is so proud of you!", config.Name),
			fmt.Sprintf("good boy/girl!"),
		}
	}

	idx := rand.Intn(len(config.Dialogs))

	fmt.Printf("%v\n", config.Dialogs[idx])
}
