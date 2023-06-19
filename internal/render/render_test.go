package render

import (
	"net/http"
	"testing"

	"github.com/Alinoureddine1/ZenStay/internal/models"
)

func TestNewTemplates(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "")
	result := AddDefaultData(&td, r)
	if result.Flash != "" {
		t.Error("Flash message not found in session when it should be")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}
