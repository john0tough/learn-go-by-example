package concurrency

type WebSiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	resultChannel := make(chan result)
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(url)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
