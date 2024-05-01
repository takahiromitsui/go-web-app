package render

import (
	"net/http"
	"testing"

	"github.com/takahiromitsui/go-web-app/internal/models"
)


func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("expected 123 but got", result.Flash)
	}
}

func getSession()(*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx :=r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateToCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplatesCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww myWriter
	
	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}
	err = RenderTemplate(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that does not exist")
	}
}