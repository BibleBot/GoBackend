package decryption

import (
	"bytes"
	"errors"
	"io/ioutil"

	"internal.kerygma.digital/kerygma-digital/biblebot/backend/utils/logger"

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
		return nil, logger.LogWithError("decrypt@read", "failed to read buffer", err)
	}

	plaintext, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, logger.LogWithError("decrypt@read", "failed to convert buffer to plaintext", err)
	}

	return plaintext, nil
}
