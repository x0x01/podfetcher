package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gregf/podfetcher/src/helpers"
)

// Setup writes a default configuration file.
func Setup() {
	configFile := helpers.ConfigPath()

	fmt.Println("\nNo configuration file found, auto generating a default config")
	fmt.Println("New config file can be found at", configFile)

	err := os.MkdirAll(filepath.Dir(configFile), 0755)
	if err != nil {
		log.Fatalf("mkdir failed %s\n", err)
	}

	file, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	config := append([]string{
		"main:\n",
		"  download: ~/podcasts\n",
		"  youtube-quality: best\n",
		"  episodes: 10\n",
		"  notify-program: notify-send\n",
		"filters:\n"})

	for _, line := range config {
		_, err = file.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
	}
}
