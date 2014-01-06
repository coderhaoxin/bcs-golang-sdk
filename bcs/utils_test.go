package bcs

import (
	"testing"
)

func TestUtils(t *testing.T) {
	if "a%3D1%26b%3D2" != UrlEncode("a=1&b=2") {
		t.Error("UrlEncode failed.")
	}

	if "YT0xJmI9Mg==" != Base64Encode([]byte("a=1&b=2")) {
		t.Error("Base64Encode failed.")
	}

}
