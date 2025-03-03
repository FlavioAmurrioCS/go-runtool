package cmd

import (
	"net/url"
	"testing"
)

func TestEnsureFullUrl(t *testing.T) {
	raw_url := "https://github.com/junegunn/fzf/releases/expanded_assets/v0.59.0"
	base_url, _ := url.Parse(raw_url)
	f := ensureFullUrl(base_url)

	actual := f("/junegunn/fzf/releases/download/v0.60.0/fzf-0.60.0-linux_loong64.tar.gz")
	expected := "https://github.com/junegunn/fzf/releases/download/v0.60.0/fzf-0.60.0-linux_loong64.tar.gz"

	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}

	actual = f("/fzf-0.60.0-linux_loong64.tar.gz")
	expected = "https://github.com/fzf-0.60.0-linux_loong64.tar.gz"
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}

	actual = f("fzf-0.60.0-linux_loong64.tar.gz")
	expected = "https://github.com/junegunn/fzf/releases/expanded_assets/v0.59.0/fzf-0.60.0-linux_loong64.tar.gz"
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
