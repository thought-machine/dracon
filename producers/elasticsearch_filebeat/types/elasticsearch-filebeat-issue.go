package types

// ElasticSearchFilebeatResult represents how a filebeat result appears in elasticsearch
type ElasticSearchFilebeatResult struct {
	Hits         elasticSearchFilebeatHits         `json:"hits"`
	Aggregations elasticSearchFilebeatAggregations `json:"aggregations"`
}

type elasticSearchFilebeatHits struct {
	Hits []elasticSearchFilebeatHit `json:"hits"`
}

type elasticSearchFilebeatHit struct {
	ID     string                      `json:"_id"`
	Source elasticSearchFilebeatSource `json:"_source"`
}

type elasticSearchFilebeatSource struct {
	Message   string                    `json:"message"`
	Timestamp string                    `json:"@timestamp"`
	Host      elasticSearchFilebeatHost `json:"host"`
}

type elasticSearchFilebeatHost struct {
	Name string `json:"name"`
}

// Note that the aggregation field name depends on the request.
type elasticSearchFilebeatAggregations struct {
	Aggregation elasticSearchFilebeatAggregation `json:"aggregation"`
}

type elasticSearchFilebeatAggregation struct {
	Buckets []elasticSearchFilebeatBucket `json:"buckets"`
}

// Note that the metric field name depends on the request.
type elasticSearchFilebeatBucket struct {
	Name   string                      `json:"key"`
	Metric elasticSearchFilebeatMetric `json:"metric"`
}

type elasticSearchFilebeatMetric struct {
	Hits elasticSearchFilebeatHits `json:"hits"`
}
