package model

import "time"

// Category represents a blog post category
type Category struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Slug        string    `db:"slug" json:"slug"`
	Description *string   `db:"description" json:"description,omitempty"`
	SortOrder   int       `db:"sort_order" json:"sort_order"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`

	// Computed fields (not stored in database)
	PostCount int `db:"-" json:"post_count,omitempty"`
}

// CreateCategoryRequest represents the request body for creating a category
type CreateCategoryRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=100"`
	Slug        string  `json:"slug" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,max=500"`
	SortOrder   int     `json:"sort_order" validate:"omitempty,min=0"`
}

// UpdateCategoryRequest represents the request body for updating a category
type UpdateCategoryRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=1,max=100"`
	Slug        *string `json:"slug" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,max=500"`
	SortOrder   *int    `json:"sort_order" validate:"omitempty,min=0"`
}

// CategoryResponse represents the response body for category endpoints
type CategoryResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description,omitempty"`
	SortOrder   int     `json:"sort_order"`
	PostCount   int     `json:"post_count"`
}

// ToResponse converts a Category to CategoryResponse
func (c *Category) ToResponse() *CategoryResponse {
	return &CategoryResponse{
		ID:          c.ID,
		Name:        c.Name,
		Slug:        c.Slug,
		Description: c.Description,
		SortOrder:   c.SortOrder,
		PostCount:   c.PostCount,
	}
}
