package decompression

import (
	"errors"
	"io"

	"archive/tar"

	"github.com/klauspost/compress/zstd"
)

// Decompress a .tar.zst file.
func Decompress(r io.Reader, w io.Writer) error {
	zstdReader, err := zstd.NewReader(r)
	if err != nil {
		return errors.New("couldn't create zstd reader")
	}
	defer zstdReader.Close()

	tarReader := tar.NewReader(zstdReader)

	_, err = io.Copy(w, tarReader)
	return err
}
