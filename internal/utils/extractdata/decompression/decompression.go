package decompression

import (
	"errors"
	"io"

	"github.com/klauspost/compress/zstd"
)

// DecompressZstd does what it's named.
func DecompressZstd(r io.Reader, w io.Writer) error {
	zstdReader, err := zstd.NewReader(r)
	if err != nil {
		return errors.New("couldn't create zstd reader")
	}
	defer zstdReader.Close()

	_, err = io.Copy(w, zstdReader)
	return err
}
