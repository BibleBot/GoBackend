package extractdata

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"backend/internal/utils/extractdata/decompression"
	"backend/internal/utils/extractdata/decryption"
	"backend/internal/utils/logger"
)

// ExtractData extracts *.tar.zst and *.tar.zst.gpg files.
func ExtractData(inputPath string, outputPath string, password string) error {
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

	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		logger.Log("err", "extractdata@write", err.Error())
		return err
	}

	return decompression.DecompressZstd(bytesReader, outputFile)
}
