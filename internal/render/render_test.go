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
	result := AddDefaultData(&td, r)
	if result == nil {
		t.Error("returned nil")
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