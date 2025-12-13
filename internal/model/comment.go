package model

import "time"

// CommentStatus represents the status of a comment
type CommentStatus string

const (
	CommentStatusPending  CommentStatus = "pending"
	CommentStatusApproved CommentStatus = "approved"
	CommentStatusSpam     CommentStatus = "spam"
)

// Comment represents a comment on a post
type Comment struct {
	ID          int64         `db:"id" json:"id"`
	PostID      int64         `db:"post_id" json:"post_id"`
	ParentID    *int64        `db:"parent_id" json:"parent_id,omitempty"`
	AuthorName  string        `db:"author_name" json:"author_name"`
	AuthorEmail string        `db:"author_email" json:"author_email"`
	Content     string        `db:"content" json:"content"`
	Status      CommentStatus `db:"status" json:"status"`
	CreatedAt   time.Time     `db:"created_at" json:"created_at"`
}

// CreateCommentRequest represents the request to create a new comment
type CreateCommentRequest struct {
	PostID      int64  `json:"post_id" validate:"required"`
	ParentID    *int64 `json:"parent_id,omitempty"`
	AuthorName  string `json:"author_name" validate:"required,min=2,max=50"`
	AuthorEmail string `json:"author_email" validate:"required,email"`
	Content     string `json:"content" validate:"required,min=1,max=2000"`
}

// UpdateCommentStatusRequest represents the request to update comment status
type UpdateCommentStatusRequest struct {
	Status CommentStatus `json:"status" validate:"required,oneof=pending approved spam"`
}

// CommentWithPost includes post information for admin listing
type CommentWithPost struct {
	Comment
	PostTitle string `db:"post_title" json:"post_title"`
	PostSlug  string `db:"post_slug" json:"post_slug"`
}

// CommentWithReplies includes nested replies
type CommentWithReplies struct {
	Comment
	Replies []Comment `json:"replies,omitempty"`
}
