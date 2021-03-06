package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"
)

const (
	testKey     string = "3077020101042024143f1db894a7300475a7da10687ed422aae67c921d93d9b85a72c8c9b05526a00a06082a8648ce3d030107a14403420004870be068597eee4240931347fc99ac4362de8e2e99204ef0f8f249b8ec8c9a1450dec73219671238530006e4fb6254876a200501f1ceb57f20f1b624e6c871e1"
	testPayload string = "00a94a368875d47544f26a5fc660512198a5660ef51465049d22a60a24e2d50f"
	testSig     string = "55a38cf365f007813a4e05eaa49814bddf203d947c9e54d27daf45df1fc20118464cc8872b12cd9f65b4bbb99e8f1fccb939329b38aff3d91e3d6d147387bd9a"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func (f fakeLayer) hasWalletetFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	// return utils.ToBytes(makeTestWallet().privateKey), nil
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("Wallet is Created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return false },
		}
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
	t.Run("Wallet is Restored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return true },
		}
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
}

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	// t.Log(s)
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	// privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// b, _ := x509.MarshalECPrivateKey(privKey)
	// t.Logf("%x", b)
	type test struct {
		Input string
		ok    bool
	}
	tests := []test{
		{testPayload, true},
		{"00a94a368875d47544f26a5fc660512198a5660ef51465049d2fa60ad4e2450f", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.Input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("x")
	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex.")
	}
}
