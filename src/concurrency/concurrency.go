package concurrency

type WebsiteChecker func(string) bool

// returns a map where each key is a url and
// the value is true if the url could be reached
func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	type result struct {
		string
		bool
	}

	results := make(map[string]bool, len(urls))
	resultsChannel := make(chan result)

	for _, url := range urls {
		url := url

		go func() {
			resultsChannel <- result{url, checker(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultsChannel
		results[result.string] = result.bool
	}

	return results
}
