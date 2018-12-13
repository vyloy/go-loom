// +build evm

package crypto

import (
	"os"
	"testing"

	"golang.org/x/crypto/ed25519"
)

func TestGenYubiSecp256k1Key(t *testing.T) {
	// check if yubiHsm config is given
	cfg := os.Getenv("YUBIHSM_CFG_FILE")
	if len(cfg) == 0 {
		t.Log("YubiHsm crypto testing disabled")
		return
	}

	t.Log("Generating YubiHSM private key")
	yubiPrivKey, err := GenYubiHsmPrivKey("ed25519", cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer yubiPrivKey.UnloadYubiHsmPrivKey()

	yubiPrivKey.SaveYubiHsmPrivKey(cfg)
}

func TestSignYubiSecp256k1(t *testing.T) {
	// check if yubiHsm config is given
	cfg := os.Getenv("YUBIHSM_CFG_FILE")
	if len(cfg) == 0 {
		t.Log("YubiHsm crypto testing disabled")
		return
	}

	t.Log("Loading YubiHSM private key")
	yubiPrivKey, err := LoadYubiHsmPrivKey("ed25519", cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer yubiPrivKey.UnloadYubiHsmPrivKey()

	t.Logf("LoadYubiHsmPrivKey succeeded")

	testMsg := []byte{'t', 'e', 's', 't'}
	sig, err := YubiHsmSign(testMsg, yubiPrivKey)
	if err != nil {
		t.Fatal(err)
	}

	if !ed25519.Verify(yubiPrivKey.pubKeyBytes[:], testMsg, sig) {
		t.Fatalf("Verification of signature has failed")
	}

	t.Logf("Sign/Verify using YubiHSM has been succeeded")
}
