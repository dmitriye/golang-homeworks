package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const top = 10

var regex = regexp.MustCompile(`[^а-яА-ЯёЁ\-]`)

func Top10(str string) []string {
	words := strings.Fields(str)

	freq := map[string]int{}
	for _, word := range words {
		if word == "-" {
			continue
		}
		key := strings.Trim(word, "-")
		key = strings.ToLower(key)
		key = regex.ReplaceAllString(key, "")
		freq[key]++
	}

	i := 0
	keys := make([]string, len(freq))
	for v := range freq {
		keys[i] = v
		i++
	}

	sort.Slice(keys, func(i, j int) bool {
		if freq[keys[i]] == freq[keys[j]] {
			return keys[i] < keys[j]
		}
		return freq[keys[i]] > freq[keys[j]]
	})

	if len(keys) < top {
		return keys
	}

	return keys[:top]
}
