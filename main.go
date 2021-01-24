package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"github.com/raspi/fs2util/pkg/vp"
	"hash"
	"os"
	"path"
	"sort"
	"strings"
)

var (
	// These are set with Makefile -X=main.VERSION, etc
	VERSION   = `v0.0.0`
	BUILD     = `dev`
	BUILDDATE = `0000-00-00T00:00:00+00:00`
)

const (
	AUTHOR   = `Pekka Järvinen`
	HOMEPAGE = `https://github.com/raspi/fs2util`
	YEAR     = 2021
)

var checksummers = map[string]hash.Hash{
	"sha256": sha256.New(),
	"sha1":   sha1.New(),
	"md5":    md5.New(),
	"sha512": sha512.New(),
}

func main() {

	extractArg := flag.Bool(`x`, false, `Extract files`)
	extractDirArg := flag.String(`edir`, `extracted`, `Extract directory name`)

	var csumAlgos []string
	for aname, _ := range checksummers {
		csumAlgos = append(csumAlgos, aname)
	}
	sort.Strings(csumAlgos)

	checksumArg := flag.Bool(`c`, false, `Calculate checksums inside the file`)
	csumAlgoArg := flag.String(`algo`, `sha256`, `Checksum algorithm: `+strings.Join(csumAlgos, `, `))

	infoArg := flag.Bool(`i`, false, `Display information about files inside the .vp file`)

	flag.Usage = func() {
		f := os.Args[0]
		_, _ = fmt.Fprintf(os.Stdout, `fs2util - Inspect FreeSpace 2 .VP file`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `Version %v %v %v`+"\n", VERSION, BUILD, BUILDDATE)
		_, _ = fmt.Fprintf(os.Stdout, `(c) %v %v- [ %v ]`+"\n", AUTHOR, YEAR, HOMEPAGE)
		_, _ = fmt.Fprintf(os.Stdout, "\n")

		_, _ = fmt.Fprintf(os.Stdout, "Parameters:\n")

		paramMaxLen := 0

		flag.VisitAll(func(f *flag.Flag) {
			l := len(f.Name)
			if l > paramMaxLen {
				paramMaxLen = l
			}
		})

		flag.VisitAll(func(f *flag.Flag) {
			padding := strings.Repeat(` `, paramMaxLen-len(f.Name))
			_, _ = fmt.Fprintf(os.Stdout, "  -%s%s   %s   default: %q\n", f.Name, padding, f.Usage, f.DefValue)
		})

		_, _ = fmt.Fprintf(os.Stdout, "\n")

		_, _ = fmt.Fprintf(os.Stdout, `List file checksums:`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `  %s -c <file.vp>`+"\n", f)
		_, _ = fmt.Fprintf(os.Stdout, `Extract file(s):`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `  %s -x <file.vp>`+"\n", f)
		_, _ = fmt.Fprintf(os.Stdout, "\n")
		_, _ = fmt.Fprintf(os.Stdout, `Examples:`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `  Show checksums:`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `    %s -c root_fs2.vp`+"\n", f)
		_, _ = fmt.Fprintf(os.Stdout, `  Extract files:`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `    %s -x root_fs2.vp`+"\n", f)
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if !*extractArg && !*checksumArg && !*infoArg {
		_, _ = fmt.Fprintf(os.Stderr, `missing parameter(s), see --help`)
		os.Exit(1)
	}

	filename := flag.Arg(flag.NArg() - 1)

	fi, err := os.Stat(filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `Not file: %q`, filename)
		os.Exit(1)
	}

	if !fi.Mode().IsRegular() {
		_, _ = fmt.Fprintf(os.Stderr, `Not file: %q`, fi.Name())
		os.Exit(1)
	}

	hasher, ok := checksummers[*csumAlgoArg]

	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, `Not valid algo: %q`, *csumAlgoArg)
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := vp.New(f, hasher)

	err = reader.ReadFile()
	if err != nil {
		panic(err)
	}

	// Get files
	for _, fi := range reader.GetFiles() {

		fInfo, err := f.Stat()
		if err != nil {
			panic(err)
		}

		if *infoArg {
			fmt.Printf(`source:%s file:%s/%s size:%d bytes, offset:%d date:%v`+"\n", fInfo.Name(), fi.Path, fi.FileName, fi.FileSize, fi.Offset, fi.Time.UTC())
		}

		if *extractArg {
			// Extract file
			expath, err := reader.ExtractFile(fi, path.Join(*extractDirArg, fInfo.Name()))
			if err != nil {
				panic(err)
			}
			fmt.Printf(`Extracted file %s/%s size:%d bytes, offset:%d (%v) to %s`+"\n", fi.Path, fi.FileName, fi.FileSize, fi.Offset, fi.Time.UTC(), expath)
		}

		if *checksumArg {
			// Get file checksum
			checksum, err := reader.ChecksumFile(fi)
			if err != nil {
				panic(err)
			}

			fmt.Printf(`%s  %s/%s/%s`+"\n", checksum, fInfo.Name(), fi.Path, fi.FileName)
		}
	}

}
