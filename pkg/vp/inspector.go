package vp

import (
	"encoding/binary"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

// Header
type RawHeader struct {
	Magic           [4]byte // VPVP
	Version         uint32  // 2
	DirectoryOffset uint32
	EntryCount      uint32
}

// File entry
type RawFileEntry struct {
	Offset        uint32
	Size          uint32
	FileName      [32]byte // null-terminated
	UnixTimeStamp uint32   // https://en.wikipedia.org/wiki/Unix_time
}

// File in more human readable form
type File struct {
	Path     string
	FileName string
	FileSize int
	Offset   int64
	Time     time.Time
}

func (rfe RawFileEntry) GetFileName() string {
	n := string(rfe.FileName[:])
	return n[0:strings.IndexByte(n, 0)]
}

type Inspector struct {
	f           io.ReadSeeker
	header      RawHeader
	files       []File
	checksummer hash.Hash // checksummer to be used (sha1, sha256, md5, ...)
}

func New(f io.ReadSeeker, checksummer hash.Hash) Inspector {
	return Inspector{
		f:           f,
		checksummer: checksummer,
	}
}

func (i *Inspector) ReadFile() error {
	err := binary.Read(i.f, binary.LittleEndian, &i.header)
	if err != nil {
		return err
	}

	// Check header
	hdr := string(i.header.Magic[:])
	if hdr != `VPVP` {
		return fmt.Errorf(`invalid header: %q`, hdr)
	}

	// Check version
	if i.header.Version != 2 {
		return fmt.Errorf(`invalid version: %d`, i.header.Version)
	}

	return i.readDirectoryIndex()
}

func (i *Inspector) readDirectoryIndex() error {
	// Go to Directory Index offset
	_, err := i.f.Seek(int64(i.header.DirectoryOffset), io.SeekStart)
	if err != nil {
		return err
	}

	fpath := `.`

	// Read directory index file entries
	for idx := 0; idx < int(i.header.EntryCount); idx++ {
		var fe RawFileEntry
		err = binary.Read(i.f, binary.LittleEndian, &fe)
		if err != nil {
			panic(err)
		}

		if fe.Size == 0 {
			// Size 0 = directory
			fpath = path.Join(fpath, fe.GetFileName())
			continue
		}

		i.files = append(i.files, File{
			Path:     fpath,
			FileName: fe.GetFileName(),
			FileSize: int(fe.Size),
			Offset:   int64(fe.Offset),
			Time:     time.Unix(int64(fe.UnixTimeStamp), 0),
		})
	}

	return nil

}

func (i Inspector) GetFiles() []File {
	return i.files
}

// Extract file
func (i Inspector) ExtractFile(target File, extractionDir string) (string, error) {
	_, err := i.f.Seek(target.Offset, io.SeekStart)
	if err != nil {
		return ``, err
	}

	buffer := make([]byte, target.FileSize)

	_, err = i.f.Read(buffer)
	if err != nil {
		return ``, err
	}

	// Write to temporary file
	err = os.MkdirAll(extractionDir, os.ModePerm)
	if err != nil {
		return ``, err
	}
	f, err := ioutil.TempFile(extractionDir, `extracted-`)
	if err != nil {
		return ``, err
	}

	// Write contents to file
	_, err = f.Write(buffer)
	if err != nil {
		return ``, err
	}

	err = f.Close()
	if err != nil {
		return ``, err
	}

	newPath := path.Join(extractionDir, target.Path)
	err = os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return ``, err
	}

	newFilePath := path.Join(newPath, target.FileName)

	// Rename temporary file to proper filename
	err = os.Rename(f.Name(), newFilePath)
	if err != nil {
		return ``, err
	}

	return newFilePath, nil
}

// Calculate checksum
func (i Inspector) ChecksumFile(target File) (string, error) {
	_, err := i.f.Seek(target.Offset, io.SeekStart)
	if err != nil {
		return ``, err
	}

	buffer := make([]byte, target.FileSize)

	_, err = i.f.Read(buffer)
	if err != nil {
		return ``, err
	}

	i.checksummer.Reset()
	_, err = i.checksummer.Write(buffer)
	if err != nil {
		return ``, err
	}

	return fmt.Sprintf(`%x`, i.checksummer.Sum(nil)), nil
}
