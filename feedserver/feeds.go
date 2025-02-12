package feedserver

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/strideynet/bsky-furry-feed/bluesky"
	"github.com/strideynet/bsky-furry-feed/store"
)

var (
	furryNewFeed     = "furry-new"
	furryHotFeed     = "furry-hot"
	furryTestFeed    = "furry-test"
	furryFursuitFeed = "furry-fursuit"
)

func getFurryHotFeed(
	ctx context.Context, queries *store.Queries, cursor string, limit int,
) ([]store.CandidatePost, error) {
	params := store.GetFurryHotFeedParams{
		Limit:         int32(limit),
		LikeThreshold: int32(4),
	}
	if cursor != "" {
		cursorTime, err := bluesky.ParseTime(cursor)
		if err != nil {
			return nil, fmt.Errorf("parsing cursor: %w", err)
		}
		params.CursorTimestamp = pgtype.Timestamptz{
			Valid: true,
			Time:  cursorTime,
		}
	}

	posts, err := queries.GetFurryHotFeed(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("executing sql: %w", err)
	}
	return posts, nil
}

func getFurryNewFeed(
	ctx context.Context, queries *store.Queries, cursor string, limit int,
) ([]store.CandidatePost, error) {
	params := store.GetFurryNewFeedParams{
		Limit: int32(limit),
	}
	if cursor != "" {
		cursorTime, err := bluesky.ParseTime(cursor)
		if err != nil {
			return nil, fmt.Errorf("parsing cursor: %w", err)
		}
		params.CursorTimestamp = pgtype.Timestamptz{
			Valid: true,
			Time:  cursorTime,
		}
	}

	posts, err := queries.GetFurryNewFeed(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("executing sql: %w", err)
	}
	return posts, nil
}

func getFurryNewFeedWithTag(
	ctx context.Context, queries *store.Queries, cursor string, limit int, tag string,
) ([]store.CandidatePost, error) {
	params := store.GetFurryNewFeedWithTagParams{
		Limit: int32(limit),
		Tag:   tag,
	}
	if cursor != "" {
		cursorTime, err := bluesky.ParseTime(cursor)
		if err != nil {
			return nil, fmt.Errorf("parsing cursor: %w", err)
		}
		params.CursorTimestamp = pgtype.Timestamptz{
			Valid: true,
			Time:  cursorTime,
		}
	}

	posts, err := queries.GetFurryNewFeedWithTag(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("executing sql: %w", err)
	}
	return posts, nil
}
