package decryption

import (
	"bytes"
	"errors"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// Decrypt a .gpg file.
func Decrypt(ciphertext []byte, password []byte, packetConfig *packet.Config) (plaintext []byte, err error) {
	decbuf := bytes.NewBuffer(ciphertext)

	armorBlock, err := armor.Decode(decbuf)
	if err != nil {
		return
	}

	failed := false
	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if failed {
			return nil, errors.New("failed to decrypt")
		}

		failed = true
		return password, nil
	}

	md, err := openpgp.ReadMessage(armorBlock.Body, nil, prompt, packetConfig)
	if err != nil {
		return
	}

	plaintext, err = ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return
	}

	return
}
