package filter

import (
	"fmt"
	"strings"

	"github.com/gregf/podfetcher/Godeps/_workspace/src/github.com/spf13/viper"
)

// Run filters podcasts based on filters from config.yml
func Run(podcastTitle string, episodeTitle string) bool {
	filterName := fmt.Sprintf("filters.%s", podcastTitle)
	filters := viper.GetStringSlice(filterName)

	for _, filter := range filters {
		if strings.Contains(episodeTitle, filter) {
			return true
		}
	}
	return false
}