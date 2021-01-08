package types

type ElasticSearchFilebeatResult struct {
    Hits         ElasticSearchFilebeatHits         `json:"hits"`
    Aggregations ElasticSearchFilebeatAggregations `json:"aggregations"`
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
    Name string `json:"name"`
}

// Note that the aggregation field name depends on the request.
type ElasticSearchFilebeatAggregations struct {
    Aggregation ElasticSearchFilebeatAggregation `json:"aggregation"`
}

type ElasticSearchFilebeatAggregation struct {
    Buckets []ElasticSearchFilebeatBucket `json:"buckets"`
}

// Note that the metric field name depends on the request.
type ElasticSearchFilebeatBucket struct {
    Name   string                      `json:"key"`
    Metric ElasticSearchFilebeatMetric `json:"metric"`
}

type ElasticSearchFilebeatMetric struct {
    Hits ElasticSearchFilebeatHits `json:"hits"`
}
