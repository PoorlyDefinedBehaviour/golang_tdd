package concurrency

type WebsiteChecker func(string) bool

// returns a map where each key is a url and
// the value is true if the url could be reached
func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool, len(urls))

	for _, url := range urls {
		results[url] = checker(url)
	}

	return results
}
