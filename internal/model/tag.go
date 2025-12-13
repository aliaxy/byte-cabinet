package model

import "time"

// Tag represents a tag that can be assigned to posts
type Tag struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Slug      string    `db:"slug" json:"slug"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// PostTag represents the many-to-many relationship between posts and tags
type PostTag struct {
	PostID int64 `db:"post_id" json:"post_id"`
	TagID  int64 `db:"tag_id" json:"tag_id"`
}

// TagWithCount extends Tag with post count for display purposes
type TagWithCount struct {
	Tag
	PostCount int `db:"post_count" json:"post_count"`
}

// CreateTagRequest represents the request body for creating a tag
type CreateTagRequest struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
	Slug string `json:"slug" validate:"omitempty,min=1,max=50"`
}

// UpdateTagRequest represents the request body for updating a tag
type UpdateTagRequest struct {
	Name string `json:"name" validate:"omitempty,min=1,max=50"`
	Slug string `json:"slug" validate:"omitempty,min=1,max=50"`
}

// TagResponse represents a tag in API responses
type TagResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	PostCount int       `json:"post_count,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// ToResponse converts a Tag to TagResponse
func (t *Tag) ToResponse() *TagResponse {
	return &TagResponse{
		ID:        t.ID,
		Name:      t.Name,
		Slug:      t.Slug,
		CreatedAt: t.CreatedAt,
	}
}

// ToResponseWithCount converts a TagWithCount to TagResponse
func (t *TagWithCount) ToResponseWithCount() *TagResponse {
	return &TagResponse{
		ID:        t.ID,
		Name:      t.Name,
		Slug:      t.Slug,
		PostCount: t.PostCount,
		CreatedAt: t.CreatedAt,
	}
}
