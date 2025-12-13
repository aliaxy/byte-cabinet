# Byte Cabinet - Design Document

> **Author:** Byte Cabinet Team  
> **Created:** 2025-01  
> **Status:** Draft  
> **Version:** 1.0

---

## Table of Contents

1. [Overview](#1-overview)
2. [Goals and Non-Goals](#2-goals-and-non-goals)
3. [System Architecture](#3-system-architecture)
4. [Tech Stack](#4-tech-stack)
5. [Data Model](#5-data-model)
6. [API Design](#6-api-design)
7. [Frontend Design](#7-frontend-design)
8. [Security Considerations](#8-security-considerations)
9. [Milestones](#9-milestones)
10. [Future Considerations](#10-future-considerations)

---

## 1. Overview

### 1.1 Background

Byte Cabinet is a personal blog system designed for recording technical notes and learning experiences. It aims to provide a clean, efficient, and feature-rich platform for publishing and organizing technical content.

### 1.2 Objectives

Build a modern, performant personal blog with the following characteristics:

- **Simple**: Easy to deploy and maintain
- **Fast**: Optimized for speed and performance
- **Feature-rich**: Complete blogging features without bloat
- **Extensible**: Easy to add new features in the future

---

## 2. Goals and Non-Goals

### 2.1 Goals

| Goal | Description |
|------|-------------|
| Single-author blog | Support one admin user for content management |
| Article management | Create, edit, delete, and publish articles |
| Markdown support | Full Markdown editing with live preview |
| Code highlighting | Syntax highlighting for code blocks |
| Categories & Tags | Organize content with categories and tags |
| Search | Full-text search across articles |
| Comments | Reader engagement through comments |
| TOC generation | Auto-generate table of contents for articles |
| View statistics | Track article view counts |
| Image upload | Upload and manage images |
| Article export | Export articles as Markdown or PDF |
| Dark mode | Support light/dark theme switching |
| RSS feed | RSS subscription support |
| Responsive design | Mobile-friendly UI |

### 2.2 Non-Goals

| Non-Goal | Reason |
|----------|--------|
| Multi-user authoring | Personal blog, only one author needed |
| Social login | Keep authentication simple |
| Real-time collaboration | Not needed for single author |
| Multi-language i18n | English content only for now |
| Payment/Subscription | Not a commercial platform |

---

## 3. System Architecture

### 3.1 High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         Client (Browser)                         │
└─────────────────────────────┬───────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Nginx (Reverse Proxy)                        │
│                  - SSL termination                               │
│                  - Static file serving                           │
│                  - Gzip compression                              │
└─────────────────────────────┬───────────────────────────────────┘
                              │
              ┌───────────────┴───────────────┐
              ▼                               ▼
┌──────────────────────────┐    ┌──────────────────────────┐
│   Vue.js Frontend        │    │   Go Fiber Backend       │
│   (Static Files)         │    │   (API Server)           │
│                          │    │                          │
│   - SPA                  │    │   - RESTful API          │
│   - Vue Router           │    │   - JWT Auth             │
│   - Pinia Store          │    │   - Business Logic       │
│   - Markdown Editor      │    │   - File Upload          │
└──────────────────────────┘    └─────────────┬────────────┘
                                              │
                                              ▼
                                ┌──────────────────────────┐
                                │      SQLite Database      │
                                │                          │
                                │   - Articles             │
                                │   - Categories           │
                                │   - Tags                 │
                                │   - Comments             │
                                │   - Settings             │
                                └──────────────────────────┘
```

### 3.2 Backend Architecture (Layered)

```
┌─────────────────────────────────────────────────────────────────┐
│                        Handler Layer                             │
│              (HTTP handlers, request/response)                   │
└─────────────────────────────┬───────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                        Service Layer                             │
│                    (Business logic)                              │
└─────────────────────────────┬───────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Repository Layer                            │
│                    (Data access, SQL)                            │
└─────────────────────────────┬───────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                         Database                                 │
│                        (SQLite)                                  │
└─────────────────────────────────────────────────────────────────┘
```

### 3.3 Directory Structure

```
byte-cabinet/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── handler/
│   │   ├── auth.go              # Authentication handlers
│   │   ├── post.go              # Article handlers
│   │   ├── category.go          # Category handlers
│   │   ├── tag.go               # Tag handlers
│   │   ├── comment.go           # Comment handlers
│   │   ├── upload.go            # File upload handlers
│   │   └── stats.go             # Statistics handlers
│   ├── middleware/
│   │   ├── auth.go              # JWT authentication
│   │   ├── cors.go              # CORS configuration
│   │   └── logger.go            # Request logging
│   ├── model/
│   │   ├── user.go              # User model
│   │   ├── post.go              # Post model
│   │   ├── category.go          # Category model
│   │   ├── tag.go               # Tag model
│   │   └── comment.go           # Comment model
│   ├── repository/
│   │   ├── user.go              # User data access
│   │   ├── post.go              # Post data access
│   │   ├── category.go          # Category data access
│   │   ├── tag.go               # Tag data access
│   │   └── comment.go           # Comment data access
│   └── service/
│       ├── auth.go              # Auth business logic
│       ├── post.go              # Post business logic
│       ├── category.go          # Category business logic
│       ├── tag.go               # Tag business logic
│       ├── comment.go           # Comment business logic
│       ├── upload.go            # Upload business logic
│       └── export.go            # Export business logic
├── pkg/
│   ├── utils/
│   │   ├── hash.go              # Password hashing
│   │   ├── jwt.go               # JWT utilities
│   │   ├── slug.go              # URL slug generation
│   │   └── markdown.go          # Markdown processing
│   └── response/
│       └── response.go          # Unified API response
├── web/                         # Vue.js frontend
├── migrations/                  # Database migrations
├── uploads/                     # Uploaded files (gitignored)
├── data/                        # SQLite database (gitignored)
└── docs/                        # Documentation
```

---

## 4. Tech Stack

### 4.1 Backend

| Component | Technology | Reason |
|-----------|------------|--------|
| Language | Go 1.21+ | Performance, simplicity, single binary |
| Web Framework | Fiber v2 | Fast, Express-like, low memory |
| Database | SQLite | Zero-config, file-based, easy backup |
| SQL Library | sqlx | Lightweight, direct SQL control |
| Migration | golang-migrate | Standard migration tool |
| JWT | golang-jwt | Industry standard auth |
| Validation | go-playground/validator | Struct validation |
| Config | Viper | Flexible configuration |

### 4.2 Frontend

| Component | Technology | Reason |
|-----------|------------|--------|
| Framework | Vue.js 3 | Composition API, modern, performant |
| Build Tool | Vite | Fast development experience |
| State Management | Pinia | Official, TypeScript support |
| Router | Vue Router 4 | Official routing solution |
| UI Framework | TBD (Tailwind CSS) | Utility-first, customizable |
| Markdown Editor | md-editor-v3 | Vue 3 compatible, feature-rich |
| HTTP Client | Axios | Reliable, interceptors |
| Code Highlighting | Shiki / Prism | Beautiful syntax highlighting |

### 4.3 Infrastructure

| Component | Technology |
|-----------|------------|
| Reverse Proxy | Nginx |
| Deployment | Docker (optional) |
| CI/CD | GitHub Actions |

---

## 5. Data Model

### 5.1 Entity Relationship Diagram

```
┌──────────────┐       ┌──────────────┐       ┌──────────────┐
│     User     │       │     Post     │       │   Category   │
├──────────────┤       ├──────────────┤       ├──────────────┤
│ id           │       │ id           │       │ id           │
│ username     │       │ title        │──────▶│ name         │
│ email        │◀──────│ slug         │       │ slug         │
│ password     │       │ content      │       │ description  │
│ avatar       │       │ summary      │       │ created_at   │
│ created_at   │       │ cover_image  │       │ updated_at   │
│ updated_at   │       │ author_id    │       └──────────────┘
└──────────────┘       │ category_id  │
                       │ status       │       ┌──────────────┐
                       │ view_count   │       │     Tag      │
                       │ created_at   │       ├──────────────┤
                       │ updated_at   │       │ id           │
                       │ published_at │◀─────▶│ name         │
                       └──────────────┘       │ slug         │
                              │               │ created_at   │
                              │               └──────────────┘
                              │                     ▲
                              │               ┌─────┴──────┐
                              ▼               │  post_tags │
                       ┌──────────────┐       ├────────────┤
                       │   Comment    │       │ post_id    │
                       ├──────────────┤       │ tag_id     │
                       │ id           │       └────────────┘
                       │ post_id      │
                       │ author_name  │
                       │ author_email │
                       │ content      │
                       │ status       │
                       │ parent_id    │
                       │ created_at   │
                       └──────────────┘
```

### 5.2 Table Definitions

#### users

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | Unique identifier |
| username | TEXT | NOT NULL UNIQUE | Login username |
| email | TEXT | NOT NULL UNIQUE | Email address |
| password_hash | TEXT | NOT NULL | Bcrypt hashed password |
| display_name | TEXT | | Display name for UI |
| avatar | TEXT | | Avatar URL |
| bio | TEXT | | Short biography |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Creation time |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Last update time |

#### categories

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | Unique identifier |
| name | TEXT | NOT NULL UNIQUE | Category name |
| slug | TEXT | NOT NULL UNIQUE | URL-friendly identifier |
| description | TEXT | | Category description |
| sort_order | INTEGER | DEFAULT 0 | Display order |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Creation time |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Last update time |

#### tags

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | Unique identifier |
| name | TEXT | NOT NULL UNIQUE | Tag name |
| slug | TEXT | NOT NULL UNIQUE | URL-friendly identifier |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Creation time |

#### posts

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | Unique identifier |
| title | TEXT | NOT NULL | Article title |
| slug | TEXT | NOT NULL UNIQUE | URL-friendly identifier |
| content | TEXT | NOT NULL | Markdown content |
| summary | TEXT | | Article summary/excerpt |
| cover_image | TEXT | | Cover image URL |
| author_id | INTEGER | NOT NULL REFERENCES users(id) | Author reference |
| category_id | INTEGER | REFERENCES categories(id) | Category reference |
| status | TEXT | NOT NULL DEFAULT 'draft' | draft/published/archived |
| view_count | INTEGER | DEFAULT 0 | View counter |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Creation time |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Last update time |
| published_at | DATETIME | | Publication time |

#### post_tags

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| post_id | INTEGER | NOT NULL REFERENCES posts(id) ON DELETE CASCADE | Post reference |
| tag_id | INTEGER | NOT NULL REFERENCES tags(id) ON DELETE CASCADE | Tag reference |
| PRIMARY KEY | | (post_id, tag_id) | Composite primary key |

#### comments

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | INTEGER | PRIMARY KEY AUTOINCREMENT | Unique identifier |
| post_id | INTEGER | NOT NULL REFERENCES posts(id) ON DELETE CASCADE | Post reference |
| parent_id | INTEGER | REFERENCES comments(id) ON DELETE CASCADE | Parent comment (for replies) |
| author_name | TEXT | NOT NULL | Commenter's name |
| author_email | TEXT | NOT NULL | Commenter's email |
| content | TEXT | NOT NULL | Comment content |
| status | TEXT | NOT NULL DEFAULT 'pending' | pending/approved/spam |
| created_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Creation time |

#### settings

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| key | TEXT | PRIMARY KEY | Setting key |
| value | TEXT | | Setting value (JSON) |
| updated_at | DATETIME | NOT NULL DEFAULT CURRENT_TIMESTAMP | Last update time |

---

## 6. API Design

### 6.1 API Conventions

- **Base URL**: `/api/v1`
- **Format**: JSON
- **Authentication**: JWT Bearer token
- **Naming**: Snake_case for JSON fields

### 6.2 Response Format

#### Success Response

```json
{
  "success": true,
  "data": { ... },
  "message": "Operation successful"
}
```

#### Error Response

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Title is required"
  }
}
```

#### Paginated Response

```json
{
  "success": true,
  "data": {
    "items": [ ... ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 100,
      "total_pages": 10
    }
  }
}
```

### 6.3 API Endpoints

#### Authentication

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/auth/login` | Admin login | No |
| POST | `/auth/logout` | Logout | Yes |
| GET | `/auth/me` | Get current user info | Yes |
| PUT | `/auth/password` | Change password | Yes |

#### Posts (Public)

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/posts` | List published posts | No |
| GET | `/posts/:slug` | Get post by slug | No |
| GET | `/posts/search` | Search posts | No |

#### Posts (Admin)

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/admin/posts` | List all posts (with drafts) | Yes |
| GET | `/admin/posts/:id` | Get post by ID | Yes |
| POST | `/admin/posts` | Create new post | Yes |
| PUT | `/admin/posts/:id` | Update post | Yes |
| DELETE | `/admin/posts/:id` | Delete post | Yes |
| POST | `/admin/posts/:id/publish` | Publish post | Yes |
| POST | `/admin/posts/:id/unpublish` | Unpublish post | Yes |

#### Categories

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/categories` | List all categories | No |
| GET | `/categories/:slug` | Get category with posts | No |
| POST | `/admin/categories` | Create category | Yes |
| PUT | `/admin/categories/:id` | Update category | Yes |
| DELETE | `/admin/categories/:id` | Delete category | Yes |

#### Tags

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/tags` | List all tags | No |
| GET | `/tags/:slug` | Get tag with posts | No |
| POST | `/admin/tags` | Create tag | Yes |
| PUT | `/admin/tags/:id` | Update tag | Yes |
| DELETE | `/admin/tags/:id` | Delete tag | Yes |

#### Comments

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/posts/:slug/comments` | Get post comments | No |
| POST | `/posts/:slug/comments` | Add comment | No |
| GET | `/admin/comments` | List all comments | Yes |
| PUT | `/admin/comments/:id/approve` | Approve comment | Yes |
| PUT | `/admin/comments/:id/spam` | Mark as spam | Yes |
| DELETE | `/admin/comments/:id` | Delete comment | Yes |

#### Upload

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/admin/upload/image` | Upload image | Yes |
| GET | `/uploads/:filename` | Get uploaded file | No |
| DELETE | `/admin/upload/:filename` | Delete uploaded file | Yes |

#### Export

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/admin/posts/:id/export/markdown` | Export as Markdown | Yes |
| GET | `/admin/posts/:id/export/pdf` | Export as PDF | Yes |

#### Statistics

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/admin/stats/overview` | Dashboard statistics | Yes |
| GET | `/admin/stats/posts` | Post statistics | Yes |

#### Settings

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/settings` | Get public settings | No |
| GET | `/admin/settings` | Get all settings | Yes |
| PUT | `/admin/settings` | Update settings | Yes |

#### RSS & Sitemap

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/feed/rss` | RSS feed | No |
| GET | `/sitemap.xml` | Sitemap | No |

---

## 7. Frontend Design

### 7.1 Page Structure

#### Public Pages

| Page | Route | Description |
|------|-------|-------------|
| Home | `/` | List of recent posts |
| Post Detail | `/posts/:slug` | Single post view |
| Category | `/categories/:slug` | Posts in category |
| Tag | `/tags/:slug` | Posts with tag |
| Search | `/search` | Search results |
| About | `/about` | About page |
| Archives | `/archives` | All posts by date |

#### Admin Pages

| Page | Route | Description |
|------|-------|-------------|
| Login | `/admin/login` | Admin login |
| Dashboard | `/admin` | Overview statistics |
| Posts | `/admin/posts` | Post management |
| Post Editor | `/admin/posts/new` | Create/edit post |
| Categories | `/admin/categories` | Category management |
| Tags | `/admin/tags` | Tag management |
| Comments | `/admin/comments` | Comment moderation |
| Settings | `/admin/settings` | Site settings |

### 7.2 Component Structure

```
src/
├── components/
│   ├── common/
│   │   ├── AppHeader.vue
│   │   ├── AppFooter.vue
│   │   ├── AppSidebar.vue
│   │   ├── Pagination.vue
│   │   ├── Loading.vue
│   │   └── ThemeToggle.vue
│   ├── post/
│   │   ├── PostCard.vue
│   │   ├── PostList.vue
│   │   ├── PostContent.vue
│   │   ├── PostToc.vue
│   │   └── PostMeta.vue
│   ├── comment/
│   │   ├── CommentList.vue
│   │   ├── CommentItem.vue
│   │   └── CommentForm.vue
│   └── admin/
│       ├── AdminLayout.vue
│       ├── MarkdownEditor.vue
│       ├── ImageUploader.vue
│       └── StatsCard.vue
├── views/
│   ├── public/
│   │   ├── HomeView.vue
│   │   ├── PostView.vue
│   │   ├── CategoryView.vue
│   │   ├── TagView.vue
│   │   ├── SearchView.vue
│   │   ├── AboutView.vue
│   │   └── ArchivesView.vue
│   └── admin/
│       ├── LoginView.vue
│       ├── DashboardView.vue
│       ├── PostsView.vue
│       ├── PostEditorView.vue
│       ├── CategoriesView.vue
│       ├── TagsView.vue
│       ├── CommentsView.vue
│       └── SettingsView.vue
├── stores/
│   ├── auth.ts
│   ├── posts.ts
│   └── settings.ts
├── api/
│   ├── client.ts
│   ├── auth.ts
│   ├── posts.ts
│   ├── categories.ts
│   ├── tags.ts
│   └── comments.ts
└── utils/
    ├── markdown.ts
    ├── date.ts
    └── storage.ts
```

---

## 8. Security Considerations

### 8.1 Authentication & Authorization

| Measure | Implementation |
|---------|----------------|
| Password hashing | bcrypt with cost factor 12 |
| JWT tokens | Short-lived access tokens (15min) |
| Refresh tokens | Stored securely, rotated on use |
| Rate limiting | Login attempts limited (5/min) |

### 8.2 Input Validation

| Measure | Implementation |
|---------|----------------|
| Input sanitization | Validate all user inputs server-side |
| XSS prevention | Sanitize HTML output, use CSP headers |
| SQL injection | Use parameterized queries (sqlx) |
| File upload | Validate file type, size limits, rename files |

### 8.3 HTTP Security Headers

```
Content-Security-Policy: default-src 'self'
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block
Strict-Transport-Security: max-age=31536000; includeSubDomains
```

### 8.4 Comment Spam Prevention

| Measure | Implementation |
|---------|----------------|
| Moderation | Comments require approval by default |
| Rate limiting | Limit comments per IP |
| Honeypot field | Hidden field to catch bots |

---

## 9. Milestones

### Phase 1: Foundation (Week 1-2)

- [ ] Project structure setup
- [ ] Database schema and migrations
- [ ] Basic configuration management
- [ ] User authentication (login/logout)
- [ ] JWT middleware

### Phase 2: Core Backend (Week 3-4)

- [ ] Post CRUD operations
- [ ] Category management
- [ ] Tag management
- [ ] Image upload
- [ ] Search functionality

### Phase 3: Frontend Basics (Week 5-6)

- [ ] Vue project setup with Vite
- [ ] Routing and layouts
- [ ] Public pages (home, post, category, tag)
- [ ] Responsive design
- [ ] Dark mode

### Phase 4: Admin Panel (Week 7-8)

- [ ] Admin authentication
- [ ] Dashboard with statistics
- [ ] Post editor with Markdown
- [ ] Category/Tag management UI
- [ ] Image upload component

### Phase 5: Advanced Features (Week 9-10)

- [ ] Comment system
- [ ] RSS feed
- [ ] Article export (Markdown/PDF)
- [ ] SEO optimization
- [ ] Sitemap generation

### Phase 6: Polish & Deploy (Week 11-12)

- [ ] Performance optimization
- [ ] Error handling improvements
- [ ] Documentation
- [ ] Docker setup
- [ ] Deployment guide

---

## 10. Future Considerations

Features that may be added in future versions:

| Feature | Description | Priority |
|---------|-------------|----------|
| Newsletter | Email subscription for new posts | Medium |
| Reading time | Estimated reading time for posts | Low |
| Series | Group related posts into series | Medium |
| Code playground | Interactive code execution | Low |
| Analytics integration | Google Analytics / Plausible | Low |
| Backup system | Automated database backups | Medium |
| Import/Export | Import from other platforms | Low |
| WebSocket | Real-time comment updates | Low |
| PWA support | Offline reading capability | Low |
| Multi-language | Content in multiple languages | Low |

---

## Appendix A: Configuration Example

```yaml
# config.yaml
server:
  port: 3000
  host: "0.0.0.0"
  mode: "development"  # development | production

database:
  driver: "sqlite"
  path: "./data/byte-cabinet.db"

jwt:
  secret: "${JWT_SECRET}"
  access_token_ttl: "15m"
  refresh_token_ttl: "7d"

upload:
  path: "./uploads"
  max_size: 10485760  # 10MB
  allowed_types:
    - "image/jpeg"
    - "image/png"
    - "image/gif"
    - "image/webp"

blog:
  title: "Byte Cabinet"
  description: "Personal technical notes and learning experiences"
  author: "Your Name"
  url: "https://your-domain.com"
  posts_per_page: 10
```

---

## Appendix B: Error Codes

| Code | HTTP Status | Description |
|------|-------------|-------------|
| UNAUTHORIZED | 401 | Authentication required |
| FORBIDDEN | 403 | Access denied |
| NOT_FOUND | 404 | Resource not found |
| VALIDATION_ERROR | 400 | Input validation failed |
| DUPLICATE_ENTRY | 409 | Resource already exists |
| INTERNAL_ERROR | 500 | Server error |
| RATE_LIMITED | 429 | Too many requests |

---

*This document will be updated as the project evolves.*