package utils

import(
	"testing"
)

func TestAES(t *testing.T) {
	origin := "Hello"
	crypted := AesEncode(origin)
	t.Logf("encode result:%v", crypted)
	result := AesDecode(crypted)
	if result != origin {
		t.Errorf("test aes failed. result:%v", result)
	}
}
