package model

import (
	"database/sql"
	"time"
)

// PostStatus represents the publication status of a post
type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
	PostStatusArchived  PostStatus = "archived"
)

// Post represents a blog post/article
type Post struct {
	ID          int64          `db:"id" json:"id"`
	Title       string         `db:"title" json:"title"`
	Slug        string         `db:"slug" json:"slug"`
	Content     string         `db:"content" json:"content"`
	Summary     sql.NullString `db:"summary" json:"summary,omitempty"`
	CoverImage  sql.NullString `db:"cover_image" json:"cover_image,omitempty"`
	AuthorID    int64          `db:"author_id" json:"author_id"`
	CategoryID  sql.NullInt64  `db:"category_id" json:"category_id,omitempty"`
	Status      PostStatus     `db:"status" json:"status"`
	ViewCount   int64          `db:"view_count" json:"view_count"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	PublishedAt sql.NullTime   `db:"published_at" json:"published_at,omitempty"`

	// Relationships (not stored in posts table)
	Author   *User     `db:"-" json:"author,omitempty"`
	Category *Category `db:"-" json:"category,omitempty"`
	Tags     []Tag     `db:"-" json:"tags,omitempty"`
}

// PostCreateRequest represents the request body for creating a post
type PostCreateRequest struct {
	Title      string  `json:"title" validate:"required,min=1,max=200"`
	Content    string  `json:"content" validate:"required"`
	Summary    string  `json:"summary" validate:"max=500"`
	CoverImage string  `json:"cover_image"`
	CategoryID *int64  `json:"category_id"`
	TagIDs     []int64 `json:"tag_ids"`
	Status     string  `json:"status" validate:"omitempty,oneof=draft published"`
}

// PostUpdateRequest represents the request body for updating a post
type PostUpdateRequest struct {
	Title      *string `json:"title" validate:"omitempty,min=1,max=200"`
	Content    *string `json:"content"`
	Summary    *string `json:"summary" validate:"omitempty,max=500"`
	CoverImage *string `json:"cover_image"`
	CategoryID *int64  `json:"category_id"`
	TagIDs     []int64 `json:"tag_ids"`
	Status     *string `json:"status" validate:"omitempty,oneof=draft published archived"`
}

// PostListQuery represents query parameters for listing posts
type PostListQuery struct {
	Page       int    `query:"page"`
	PageSize   int    `query:"page_size"`
	Status     string `query:"status"`
	CategoryID *int64 `query:"category_id"`
	TagID      *int64 `query:"tag_id"`
	Search     string `query:"search"`
	OrderBy    string `query:"order_by"` // created_at, updated_at, published_at, view_count
	Order      string `query:"order"`    // asc, desc
}

// PostResponse represents a post in API responses
type PostResponse struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Slug        string            `json:"slug"`
	Content     string            `json:"content,omitempty"`
	Summary     string            `json:"summary,omitempty"`
	CoverImage  string            `json:"cover_image,omitempty"`
	Author      *UserResponse     `json:"author,omitempty"`
	Category    *CategoryResponse `json:"category,omitempty"`
	Tags        []TagResponse     `json:"tags,omitempty"`
	Status      string            `json:"status"`
	ViewCount   int64             `json:"view_count"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	PublishedAt *time.Time        `json:"published_at,omitempty"`
}

// ToResponse converts a Post to PostResponse
func (p *Post) ToResponse() *PostResponse {
	resp := &PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Slug:      p.Slug,
		Content:   p.Content,
		Status:    string(p.Status),
		ViewCount: p.ViewCount,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	if p.Summary.Valid {
		resp.Summary = p.Summary.String
	}

	if p.CoverImage.Valid {
		resp.CoverImage = p.CoverImage.String
	}

	if p.PublishedAt.Valid {
		resp.PublishedAt = &p.PublishedAt.Time
	}

	if p.Author != nil {
		resp.Author = p.Author.ToResponse()
	}

	if p.Category != nil {
		resp.Category = p.Category.ToResponse()
	}

	if len(p.Tags) > 0 {
		resp.Tags = make([]TagResponse, len(p.Tags))
		for i, tag := range p.Tags {
			resp.Tags[i] = *tag.ToResponse()
		}
	}

	return resp
}

// IsPublished returns true if the post is published
func (p *Post) IsPublished() bool {
	return p.Status == PostStatusPublished
}

// IsDraft returns true if the post is a draft
func (p *Post) IsDraft() bool {
	return p.Status == PostStatusDraft
}
