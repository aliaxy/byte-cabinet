# Contributing Guide

Thank you for your interest in contributing to **Byte Cabinet**! This document outlines the commit conventions and guidelines for this project.

## Commit Message Convention

We use **emoji-prefixed** commit messages to make the git history more readable and easier to navigate. Each commit message should start with an emoji that represents the type of change.

### Commit Format

```
<emoji> <type>: <short description>

[optional body]

[optional footer]
```

### Emoji Reference

| Emoji | Code | Type | Description |
|-------|------|------|-------------|
| ✨ | `:sparkles:` | feat | Introduce new features |
| 🐛 | `:bug:` | fix | Fix a bug |
| 📝 | `:memo:` | docs | Add or update documentation |
| 💄 | `:lipstick:` | style | Add or update UI and style files |
| ♻️ | `:recycle:` | refactor | Refactor code |
| ⚡️ | `:zap:` | perf | Improve performance |
| ✅ | `:white_check_mark:` | test | Add, update, or pass tests |
| 🔧 | `:wrench:` | config | Add or update configuration files |
| 🏗️ | `:building_construction:` | arch | Make architectural changes |
| 📦 | `:package:` | build | Build system or external dependencies |
| 🚀 | `:rocket:` | deploy | Deploy stuff |
| 🔒 | `:lock:` | security | Fix security issues |
| 🗑️ | `:wastebasket:` | remove | Remove code or files |
| 🚚 | `:truck:` | move | Move or rename resources |
| 🎨 | `:art:` | format | Improve structure/format of the code |
| 🔥 | `:fire:` | prune | Remove dead code |
| 🚧 | `:construction:` | wip | Work in progress |
| 💚 | `:green_heart:` | ci | Fix CI build |
| ⬆️ | `:arrow_up:` | upgrade | Upgrade dependencies |
| ⬇️ | `:arrow_down:` | downgrade | Downgrade dependencies |
| 🎉 | `:tada:` | init | Begin a project |
| 🔀 | `:twisted_rightwards_arrows:` | merge | Merge branches |
| ⏪ | `:rewind:` | revert | Revert changes |
| 🐳 | `:whale:` | docker | Docker related changes |
| 🗃️ | `:card_file_box:` | database | Database related changes |
| 🌐 | `:globe_with_meridians:` | i18n | Internationalization and localization |
| 💡 | `:bulb:` | comment | Add or update comments in source code |
| 🍱 | `:bento:` | assets | Add or update assets |
| 🙈 | `:see_no_evil:` | gitignore | Add or update .gitignore |

### Examples

```
🎉 init: initialize project with Go Fiber and Vue

✨ feat: add user authentication module

🐛 fix: resolve login redirect issue

📝 docs: update API documentation

♻️ refactor: simplify database connection logic

⚡️ perf: optimize image loading with lazy load

✅ test: add unit tests for auth service

🔧 config: update nginx configuration

🗃️ database: add migration for posts table
```

## Branch Naming Convention

Use descriptive branch names with the following format:

```
<type>/<short-description>
```

Examples:
- `feat/user-authentication`
- `fix/login-redirect`
- `docs/api-documentation`
- `refactor/database-layer`

## Code Style

- Follow Go best practices and use `gofmt` for formatting
- Follow Vue.js style guide for frontend code
- Use ESLint and Prettier for JavaScript/TypeScript formatting
- Write meaningful comments in English

## Pull Request Process

1. Create a feature branch from `main`
2. Make your changes following the conventions above
3. Ensure all tests pass
4. Submit a pull request with a clear description
5. Wait for code review

## Questions?

If you have any questions, feel free to open an issue for discussion.