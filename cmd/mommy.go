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

const (
	NotFound = iota
	Permission
	TimeOut // probably never happens at our context
	Unknown
)

type PathError struct {
	err  int
	path string
}

func (e PathError) Error() string {
	switch e.err {
	case NotFound:
		return fmt.Sprintf("(file %v) %v", e.path, "not found")
	case Permission:
		return fmt.Sprintf("(file %v) %v", e.path, "permission denied")
	case TimeOut:
		return fmt.Sprintf("(file %v) %v", e.path, "time")
	default:
		return "unknown error"
	}
}

const Mommy = "mommy"

func get_config_path() (string, error) {
	var config_dir string

	home_path := os.Getenv("HOME")
	config_home_path := os.Getenv("XDG_CONFIG_HOME")

	if config_home_path != "" {
		config_dir = fmt.Sprintf("%v/%v", config_home_path, Mommy)
	} else {
		if home_path != "" {
			config_dir = fmt.Sprintf("%v/.config/%v", home_path, Mommy)
		} else {
			return "", PathError{err: NotFound}
		}
	}

	if err := os.Mkdir(config_dir, 0o750); err != nil && !os.IsExist(err) {
		if os.IsPermission(err) {
			return "", PathError{err: Permission}
		} else {
			return "", PathError{err: Unknown}
		}
	}

	var config_path string = fmt.Sprintf("%v/%v", config_dir, "config.json")

	return config_path, nil
}

func main() {

	x := os.Getenv("?")
	fmt.Println(x)

	config_path, err := get_config_path()
	if err != nil {
		log.Fatal(err)
	}

	config_blob, err := os.ReadFile(config_path)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(config_path)
			if err := os.WriteFile(config_path, []byte("{}\n"), 0o660); err != nil {
				log.Fatal(err)
			}
			config_blob = []byte("{}")
		} else {
			log.Fatal(err)
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
