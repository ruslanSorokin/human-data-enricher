//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/incu6us/goimports-reviser/v3"
	_ "github.com/segmentio/golines"
	_ "mvdan.cc/gofumpt"
)
