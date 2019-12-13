package enrichment

import (
	"context"
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/thought-machine/dracon/pkg/enrichment/db"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

// GetHash returns the hash of an issue
func GetHash(i *v1.Issue) string {
	h := md5.New()
	io.WriteString(h, i.GetTarget())
	io.WriteString(h, i.GetType())
	io.WriteString(h, i.GetTitle())
	io.WriteString(h, i.GetSource())
	io.WriteString(h, i.GetSeverity().String())
	io.WriteString(h, fmt.Sprintf("%f", i.GetCvss()))
	io.WriteString(h, i.GetConfidence().String())
	io.WriteString(h, i.GetDescription())

	return fmt.Sprintf("%x", h.Sum(nil))
}

// NewEnrichedIssue returns a new enriched issue from a raw issue
func NewEnrichedIssue(i *v1.Issue) *v1.EnrichedIssue {
	return &v1.EnrichedIssue{
		RawIssue:      i,
		FirstSeen:     ptypes.TimestampNow(),
		Count:         1,
		FalsePositive: false,
		UpdatedAt:     ptypes.TimestampNow(),
		Hash:          GetHash(i),
	}
}

// UpdateEnrichedIssue updates a given enriched issue
func UpdateEnrichedIssue(i *v1.EnrichedIssue) {
	i.Count++
	i.UpdatedAt = ptypes.TimestampNow()
}

// EnrichIssue enriches a given issue, returning an enriched issue once processed
func EnrichIssue(db *db.DB, i *v1.Issue) (*v1.EnrichedIssue, error) {
	hash := GetHash(i)
	enrichedIssue, err := db.GetIssueByHash(hash)
	if errors.Is(err, sql.ErrNoRows) {
		// create issue
		enrichedIssue = NewEnrichedIssue(i)
		err := db.CreateIssue(context.Background(), enrichedIssue)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return enrichedIssue, nil
	} else if err != nil {
		return nil, err
	}
	// update issue
	UpdateEnrichedIssue(enrichedIssue)
	if err := db.UpdateIssue(context.Background(), enrichedIssue); err != nil {
		return nil, err
	}
	return enrichedIssue, nil
}
