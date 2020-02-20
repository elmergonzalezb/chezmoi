//go:generate go run ./internal/generate-assets -o cmd/docs.gen.go -tags=!noembeddocs docs/CHANGES.md docs/CONTRIBUTING.md docs/FAQ.md docs/HOWTO.md docs/INSTALL.md docs/QUICKSTART.md docs/REFERENCE.md
//go:generate go run ./internal/generate-assets -o internal/cmd/templates.gen.go assets/templates/COMMIT_MESSAGE.tmpl
//go:generate go run ./internal/generate-helps -o cmd/helps.gen.go -i docs/REFERENCE.md

package main

import "github.com/twpayne/chezmoi/cmd"

func main() {
	cmd.Execute()
}
