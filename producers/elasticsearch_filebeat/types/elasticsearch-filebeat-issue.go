package types

type ElasticSearchFilebeatResult struct {
	Hits ElasticSearchFilebeatHits `json:"hits"`
}

type ElasticSearchFilebeatHits struct {
	Hits []ElasticSearchFilebeatHit `json:"hits"`
}

type ElasticSearchFilebeatHit struct {
	ID        string                      `json:"_id"`
	Source    ElasticSearchFilebeatSource `json:"_source"`
}

type ElasticSearchFilebeatSource struct {
	Message   string                    `json:"message"`
	Timestamp string                    `json:"@timestamp"`
	Host      ElasticSearchFilebeatHost `json:"host"`
}

type ElasticSearchFilebeatHost struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
}
