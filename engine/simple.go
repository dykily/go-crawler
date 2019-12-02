package engine

import (
	"../fetcher"
	"log"
)

type SimpleEngine struct {

}

func Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		}

		ParseResult := r.ParserFunc(body)
		requests = append(requests, ParseResult.Requests...)

		for _, item := range ParseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s ", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}