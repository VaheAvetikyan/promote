package configuration

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Setenv("CONFIG_FILE_NAME", "promotions.config")
	t.Setenv("CONFIG_FILE_PATH", "./")
	tests := struct {
		name    string
		want    *Constants
		wantErr bool
	}{
		"a", &Constants{HOST: "localhost", PORT: "1321"}, false,
	}
	result, err := New()
	if err != nil {
		t.Error("Error occurred")
	}
	if result.PORT != tests.want.PORT {
		t.Errorf("got %s, want %s", result.PORT, tests.want.PORT)
	}
	if result.HOST != "localhost" {
		t.Errorf("got %s, want %s", result.PORT, tests.want.PORT)
	}
}
