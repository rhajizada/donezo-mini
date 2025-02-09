// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tags.sql

package repository

import (
	"context"
)

const addTagToItemByID = `-- name: AddTagToItemByID :exec
INSERT INTO tags (item_id, tag)
VALUES (?, ?)
ON CONFLICT(item_id, tag) DO NOTHING
`

type AddTagToItemByIDParams struct {
	ItemID int64  `json:"itemId"`
	Tag    string `json:"tag"`
}

func (q *Queries) AddTagToItemByID(ctx context.Context, arg AddTagToItemByIDParams) error {
	_, err := q.db.ExecContext(ctx, addTagToItemByID, arg.ItemID, arg.Tag)
	return err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM tags WHERE tag = ?
`

func (q *Queries) DeleteTag(ctx context.Context, tag string) error {
	_, err := q.db.ExecContext(ctx, deleteTag, tag)
	return err
}

const listItemsByTag = `-- name: ListItemsByTag :many
SELECT 
    i.id, i.board_id, i.title, i.description, i.completed, 
    i.created_at, i.last_updated_at
FROM items i
JOIN tags t ON i.id = t.item_id
WHERE t.tag = ?
ORDER BY i.created_at
`

func (q *Queries) ListItemsByTag(ctx context.Context, tag string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItemsByTag, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.BoardID,
			&i.Title,
			&i.Description,
			&i.Completed,
			&i.CreatedAt,
			&i.LastUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTags = `-- name: ListTags :many
SELECT DISTINCT tag FROM tags ORDER BY tag
`

func (q *Queries) ListTags(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		items = append(items, tag)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTagsByItemID = `-- name: ListTagsByItemID :many
SELECT tag FROM tags WHERE item_id = ? ORDER BY tag
`

func (q *Queries) ListTagsByItemID(ctx context.Context, itemID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listTagsByItemID, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		items = append(items, tag)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeTagFromItemByID = `-- name: RemoveTagFromItemByID :exec
DELETE FROM tags
WHERE item_id = ? AND tag = ?
`

type RemoveTagFromItemByIDParams struct {
	ItemID int64  `json:"itemId"`
	Tag    string `json:"tag"`
}

func (q *Queries) RemoveTagFromItemByID(ctx context.Context, arg RemoveTagFromItemByIDParams) error {
	_, err := q.db.ExecContext(ctx, removeTagFromItemByID, arg.ItemID, arg.Tag)
	return err
}
