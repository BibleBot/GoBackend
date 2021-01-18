package decryption

import (
	"backend/internal/utils/logger"
	"bytes"
	"errors"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

// Decrypt a .gpg file.
func Decrypt(ciphertext []byte, password []byte, packetConfig *packet.Config) ([]byte, error) {
	decbuf := bytes.NewBuffer(ciphertext)

	failed := false
	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if failed {
			return nil, errors.New("failed to decrypt - invalid password?")
		}

		failed = true
		return password, nil
	}

	md, err := openpgp.ReadMessage(decbuf, nil, prompt, packetConfig)
	if err != nil {
		logger.Log("err", "decrypt@read", "failed to read buffer")
		return nil, err
	}

	plaintext, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		logger.Log("err", "decrypt@read", "failed to convert buffer to plaintext")
		return nil, err
	}

	return plaintext, nil
}
