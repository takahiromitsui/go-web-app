package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplatesCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
}