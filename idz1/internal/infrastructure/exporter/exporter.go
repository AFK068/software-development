package exporter

import "github.com/AFK068/bot/internal/domain"

type Exporter interface {
	domain.ExportVisitor
	Export(fileName string) error
}
