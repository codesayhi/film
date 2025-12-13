package gosimpleslug

import goslug "github.com/gosimple/slug"

type Generator struct{}

func New() *Generator { return &Generator{} }

func (g *Generator) Make(input string) string {
	return goslug.Make(input)
}
