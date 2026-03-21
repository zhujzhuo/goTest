# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Building and Running

This is a Go codebase with multiple standalone programs. Run individual files with:
```
go run <filename>.go
```

For test files:
```
go test ./...
```

## Project Structure

- **hackerrank/** - HackerRank algorithm problems (Go solutions)
- **leetcode/** - LeetCode problems with common patterns:
  - Tree algorithms (binary search trees, traversal)
  - Array/string manipulation
  - Dynamic programming
  - Two-pointer techniques
- **binarytree/** - Binary tree data structures and algorithms (traversal, sorting)
- **dos_prepare/** - Database connection pooling and prepared statement testing

## Key Patterns

- **Concurrency**: Uses both `sync.WaitGroup` and channels for goroutine coordination
- **Context**: Uses `context` package for timeout/cancellation handling
- **Database**: MySQL examples with `database/sql`, prepared statements, and transactions
- **Logging**: Uses `github.com/google/logger` for structured logging

## Dependencies

Main external packages:
- `github.com/go-sql-driver/mysql` - MySQL driver
- `github.com/google/logger` -_logging_
- `net/http/pprof` - profiling
