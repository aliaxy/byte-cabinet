package response

import (
	"github.com/gofiber/fiber/v2"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error details in a response
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Pagination represents pagination metadata
type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// PaginatedData represents paginated response data
type PaginatedData struct {
	Items      interface{} `json:"items"`
	Pagination Pagination  `json:"pagination"`
}

// Error codes
const (
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeForbidden    = "FORBIDDEN"
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeValidation   = "VALIDATION_ERROR"
	ErrCodeDuplicate    = "DUPLICATE_ENTRY"
	ErrCodeInternal     = "INTERNAL_ERROR"
	ErrCodeRateLimited  = "RATE_LIMITED"
	ErrCodeBadRequest   = "BAD_REQUEST"
	ErrCodeInvalidInput = "INVALID_INPUT"
)

// OK sends a success response with data
func OK(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Data:    data,
	})
}

// OKWithMessage sends a success response with data and message
func OKWithMessage(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Data:    data,
		Message: message,
	})
}

// Created sends a 201 response with data
func Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Success: true,
		Data:    data,
		Message: "Created successfully",
	})
}

// NoContent sends a 204 response
func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// Paginated sends a paginated response
func Paginated(c *fiber.Ctx, items interface{}, page, pageSize int, total int64) error {
	totalPages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPages++
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Data: PaginatedData{
			Items: items,
			Pagination: Pagination{
				Page:       page,
				PageSize:   pageSize,
				Total:      total,
				TotalPages: totalPages,
			},
		},
	})
}

// Error sends an error response with the given status code
func Error(c *fiber.Ctx, status int, code, message string) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	})
}

// BadRequest sends a 400 error response
func BadRequest(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusBadRequest, ErrCodeBadRequest, message)
}

// ValidationError sends a 400 validation error response
func ValidationError(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusBadRequest, ErrCodeValidation, message)
}

// Unauthorized sends a 401 error response
func Unauthorized(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Authentication required"
	}
	return Error(c, fiber.StatusUnauthorized, ErrCodeUnauthorized, message)
}

// Forbidden sends a 403 error response
func Forbidden(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Access denied"
	}
	return Error(c, fiber.StatusForbidden, ErrCodeForbidden, message)
}

// NotFound sends a 404 error response
func NotFound(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Resource not found"
	}
	return Error(c, fiber.StatusNotFound, ErrCodeNotFound, message)
}

// Conflict sends a 409 error response (for duplicate entries)
func Conflict(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusConflict, ErrCodeDuplicate, message)
}

// TooManyRequests sends a 429 error response
func TooManyRequests(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "Too many requests, please try again later"
	}
	return Error(c, fiber.StatusTooManyRequests, ErrCodeRateLimited, message)
}

// InternalError sends a 500 error response
func InternalError(c *fiber.Ctx, message string) error {
	if message == "" {
		message = "An internal error occurred"
	}
	return Error(c, fiber.StatusInternalServerError, ErrCodeInternal, message)
}

// InternalErrorWithLog sends a 500 error and logs the actual error
func InternalErrorWithLog(c *fiber.Ctx, err error) error {
	// In production, you'd want to log this error
	// log.Printf("Internal error: %v", err)
	return InternalError(c, "An internal error occurred")
}
