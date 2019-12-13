package db

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

type issue struct {
	Hash          string    `db:"hash"`
	FirstSeen     time.Time `db:"first_seen"`
	Occurrences   uint64    `db:"occurrences"`
	FalsePositive bool      `db:"false_positive"`
	UpdatedAt     time.Time `db:"updated_at"`

	Target      string  `db:"target"`
	Type        string  `db:"type"`
	Title       string  `db:"title"`
	Severity    int32   `db:"severity"`
	CVSS        float64 `db:"cvss"`
	Confidence  int32   `db:"confidence"`
	Description string  `db:"description"`
	Source      string  `db:"source"`
}

func toDBIssue(i *v1.EnrichedIssue) (*issue, error) {
	firstSeen, err := ptypes.Timestamp(i.GetFirstSeen())
	if err != nil {
		return nil, err
	}
	updatedAt, err := ptypes.Timestamp(i.GetUpdatedAt())
	if err != nil {
		return nil, err
	}
	return &issue{
		Hash:          i.GetHash(),
		FirstSeen:     firstSeen,
		Occurrences:   i.GetCount(),
		FalsePositive: i.GetFalsePositive(),
		UpdatedAt:     updatedAt,
		Target:        i.RawIssue.GetTarget(),
		Type:          i.RawIssue.GetType(),
		Title:         i.RawIssue.GetTitle(),
		Severity:      int32(i.RawIssue.GetSeverity()),
		CVSS:          i.RawIssue.GetCvss(),
		Confidence:    int32(i.RawIssue.GetConfidence()),
		Description:   i.RawIssue.GetDescription(),
		Source:        i.RawIssue.GetSource(),
	}, nil
}

func toEnrichedIssue(i *issue) (*v1.EnrichedIssue, error) {
	firstSeen, err := ptypes.TimestampProto(i.FirstSeen)
	if err != nil {
		return nil, err
	}
	return &v1.EnrichedIssue{
		Hash:          i.Hash,
		FirstSeen:     firstSeen,
		Count:         i.Occurrences,
		FalsePositive: i.FalsePositive,
		RawIssue: &v1.Issue{
			Target:      i.Target,
			Type:        i.Type,
			Title:       i.Title,
			Severity:    v1.Severity(i.Severity),
			Cvss:        i.CVSS,
			Confidence:  v1.Confidence(i.Confidence),
			Description: i.Description,
			Source:      i.Source,
		},
	}, nil
}

// GetIssueByHash returns an issue given its hash
func (db *DB) GetIssueByHash(hash string) (*v1.EnrichedIssue, error) {
	i := issue{}
	if err := db.Get(&i, `SELECT * FROM issues WHERE "hash"=$1`, hash); err != nil {
		return nil, err
	}
	return toEnrichedIssue(&i)
}

// CreateIssue creates the given enriched issue on the database
func (db *DB) CreateIssue(ctx context.Context, eI *v1.EnrichedIssue) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	i, err := toDBIssue(eI)
	if err != nil {
		return err
	}
	_, err = tx.NamedExec(`INSERT INTO
issues (
	"target",
	"type",
	"title",
	severity,
	cvss,
	confidence,
	"description",
	source,
	"hash",
	first_seen,
	occurrences,
	false_positive,
	updated_at
) VALUES (
	:target,
	:type,
	:title,
	:severity,
	:cvss,
	:confidence,
	:description,
	:source,
	:hash,
	:first_seen,
	:occurrences,
	:false_positive,
	:updated_at);`,
		map[string]interface{}{
			"target":         i.Target,
			"type":           i.Type,
			"title":          i.Title,
			"severity":       i.Severity,
			"cvss":           i.CVSS,
			"confidence":     i.Confidence,
			"description":    i.Description,
			"source":         i.Source,
			"hash":           i.Hash,
			"first_seen":     i.FirstSeen,
			"occurrences":    i.Occurrences,
			"false_positive": i.FalsePositive,
			"updated_at":     i.UpdatedAt,
		},
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// UpdateIssue updates a given enriched issue on the database
func (db *DB) UpdateIssue(ctx context.Context, eI *v1.EnrichedIssue) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	i, err := toDBIssue(eI)
	if err != nil {
		return err
	}
	_, err = tx.NamedExec(`UPDATE issues
SET
	occurrences=:occurrences,
	updated_at=:updated_at
WHERE "hash"=:hash;`,
		map[string]interface{}{
			"occurrences": i.Occurrences,
			"updated_at":  i.UpdatedAt,
			"hash":        i.Hash,
		},
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
