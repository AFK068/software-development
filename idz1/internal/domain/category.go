package domain

type Category struct {
	ID   int    `json:"id" yaml:"id" `
	Type string `json:"type" yaml:"type"`
	Name string `json:"name" yaml:"name"`
}

func (c *Category) Accept(visitor ExportVisitor) error {
	return visitor.VisitCategory(c)
}
