package internal

import (
	"log/slog"

	g "github.com/serpapi/google-search-results-golang"
)

func Search(query, serpApiKey string) (map[string]interface{}, error) {
	searchMap := make(map[string]string)
	searchMap["q"] = query
	searchQuery := g.NewGoogleSearch(searchMap, serpApiKey)
	result, err := searchQuery.GetJSON()
	if err != nil {
		return nil, err
	}
	results := result["organic_results"].([]interface{})
	first_result := results[0].(map[string]interface{})
	slog.Info("result", "Result", first_result)
	return first_result, nil
}
