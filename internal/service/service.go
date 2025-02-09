package service

import (
	"context"

	"github.com/rhajizada/donezo-mini/internal/repository"
)

type Service struct {
	Repo *repository.Queries
}

func New(r *repository.Queries) *Service {
	return &Service{
		Repo: r,
	}
}

func (s *Service) ListBoards(ctx context.Context) (*[]Board, error) {
	data, err := s.Repo.ListBoards(ctx)
	if err != nil {
		return nil, err
	}
	boards := make([]Board, len(data))
	for i, b := range data {
		boards[i] = Board{b}
	}

	return &boards, nil
}

func (s *Service) CreateBoard(ctx context.Context, boardName string) (*Board, error) {
	data, err := s.Repo.CreateBoard(ctx, boardName)
	if err != nil {
		return nil, err
	}
	return &Board{data}, nil
}

func (s *Service) UpdateBoard(ctx context.Context, board *Board) (*Board, error) {
	b := repository.UpdateBoardByIDParams{
		Name: board.Name,
		ID:   board.ID,
	}
	data, err := s.Repo.UpdateBoardByID(ctx, b)
	if err != nil {
		return nil, err
	}
	return &Board{data}, nil
}

func (s *Service) DeleteBoard(ctx context.Context, board *Board) error {
	return s.Repo.DeleteBoardByID(ctx, board.ID)
}

func (s *Service) ListItems(ctx context.Context, board *Board) (*[]Item, error) {
	data, err := s.Repo.ListItemsByBoardID(ctx, board.ID)
	if err != nil {
		return nil, err
	}
	items := make([]Item, len(data))
	for i, v := range data {
		tags, err := s.Repo.ListTagsByItemID(ctx, v.ID)
		if err != nil {
			tags = make([]string, 0)
		}
		items[i] = Item{
			v,
			tags,
		}
	}
	return &items, nil
}

func (s *Service) CreateItem(ctx context.Context, board *Board, title string, description string) (*Item, error) {
	i := repository.CreateItemParams{
		BoardID:     board.ID,
		Title:       title,
		Description: description,
	}

	data, err := s.Repo.CreateItem(ctx, i)
	if err != nil {
		return nil, err
	}
	return &Item{
		data,
		make([]string, 0),
	}, nil
}

func (s *Service) UpdateItem(ctx context.Context, item *Item) (*Item, error) {
	i := repository.UpdateItemByIDParams{
		Title:       item.Title,
		Description: item.Description,
		Completed:   item.Completed,
		ID:          item.ID,
	}

	data, err := s.Repo.UpdateItemByID(ctx, i)
	if err != nil {
		return nil, err
	}

	// Fetch existing tags from the database
	existingTags, err := s.Repo.ListTagsByItemID(ctx, data.ID)
	if err != nil {
		existingTags = make([]string, 0)
	}

	// Create maps for efficient lookup
	existingTagsMap := make(map[string]struct{}, len(existingTags))
	for _, tag := range existingTags {
		existingTagsMap[tag] = struct{}{}
	}

	newTagsMap := make(map[string]struct{}, len(item.Tags))
	for _, tag := range item.Tags {
		newTagsMap[tag] = struct{}{}
	}

	// Remove tags that are not in the updated item
	for _, tag := range existingTags {
		if _, found := newTagsMap[tag]; !found {
			err := s.Repo.RemoveTagFromItemByID(ctx, repository.RemoveTagFromItemByIDParams{
				ItemID: data.ID,
				Tag:    tag,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	// Add new tags that are missing in the database
	for _, tag := range item.Tags {
		if _, found := existingTagsMap[tag]; !found {
			err := s.Repo.AddTagToItemByID(ctx, repository.AddTagToItemByIDParams{
				ItemID: data.ID,
				Tag:    tag,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return &Item{
		data,
		item.Tags, // Return the updated tags
	}, nil
}

func (s *Service) DeleteItem(ctx context.Context, item *Item) error {
	return s.Repo.DeleteItemByID(ctx, item.ID)
}

func (s *Service) ListTags(ctx context.Context) ([]string, error) {
	return s.Repo.ListTags(ctx)
}

func (s *Service) DeleteTag(ctx context.Context, tag string) error {
	return s.Repo.DeleteTag(ctx, tag)
}

func (s *Service) CountItemsByTag(ctx context.Context, tag string) (int64, error) {
	return s.Repo.CountItemsByTag(ctx, tag)
}
