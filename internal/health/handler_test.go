package health

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakeStatusService struct{}

func (f *fakeStatusService) GetStatus() string {
	return "fake-status"
}

func TestHealthHandler(t *testing.T) {
	fakeService := &fakeStatusService{}
	handler := NewHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	recorder := httptest.NewRecorder()
	handler.Check(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", recorder.Code)
	}

	expected := `{"status":"fake-status"}`

	if strings.TrimSpace(recorder.Body.String()) != expected {
		t.Errorf("expected %s, got %s", expected, recorder.Body.String())
	}
}
