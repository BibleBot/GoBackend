package extractdata

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"

	"github.com/BibleBot/backend/internal/utils/extractdata/decompression"
	"github.com/BibleBot/backend/internal/utils/extractdata/decryption"
	"github.com/BibleBot/backend/internal/utils/logger"
)

// ExtractData extracts *.tar.zst and *.tar.zst.gpg files.
func ExtractData(inputPath string, password string) error {
	inputFile, err := ioutil.ReadFile(inputPath)
	if err != nil {
		logger.Log("err", "extractdata@read", err.Error())
		return err
	}

	var bytesReader io.Reader
	if strings.HasSuffix(inputPath, ".gpg") {
		// We can assume that we'll be decrypting the file if it ends with '.gpg'.
		decryptionKey := []byte(password)
		decryptedData, err := decryption.Decrypt(inputFile, decryptionKey, nil)
		if err != nil {
			logger.Log("err", "extractdata@decrypt", err.Error())
			return err
		}

		bytesReader = bytes.NewReader(decryptedData)
	} else {
		bytesReader = bytes.NewReader(inputFile)
	}

	return decompression.Decompress(bytesReader)
}