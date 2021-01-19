package decompression

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/zstd"
)

// Decompress decompresses a .tar.zst file.
func Decompress(r io.Reader) error {
	zstdReader, err := zstd.NewReader(r)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs("./data/usx/")
	if err != nil {
		return err
	}

	tr := tar.NewReader(zstdReader)
	for {
		hdr, err := tr.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fInfo := hdr.FileInfo()
		fileName := hdr.Name
		absFileName := filepath.Join(absPath, fileName)

		if fInfo.Mode().IsDir() {
			if err := os.MkdirAll(absFileName, 0755); err != nil {
				return err
			}

			continue
		}

		file, err := os.OpenFile(
			absFileName,
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			fInfo.Mode().Perm(),
		)

		if err != nil {
			return err
		}

		n, cpErr := io.Copy(file, tr)
		if closeErr := file.Close(); closeErr != nil {
			return err
		}

		if cpErr != nil {
			return cpErr
		}

		if n != fInfo.Size() {
			return fmt.Errorf("wrote %d, want %d", n, fInfo.Size())
		}
	}

	return nil
}
