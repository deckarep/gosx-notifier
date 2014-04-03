package gosxnotifier

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	zipPath        = "terminal-notifier.temp.zip"
	executablePath = "terminal-notifier.app/Contents/MacOS/terminal-notifier"
	tempDirSuffix  = "gosxnotifier"
)

var (
	rootPath  string
	FinalPath string
)

func init() {
	err := installTerminalNotifier()
	if err != nil {
		log.Fatal("Could not install Terminal Notifier to a temp directory")
	} else {
		FinalPath = filepath.Join(rootPath, executablePath)
	}
}

func installTerminalNotifier() error {
	rootPath = filepath.Join(os.TempDir(), tempDirSuffix)

	//if terminal-notifier.app already installed no-need to re-install
	if exists(filepath.Join(rootPath, executablePath)) {
		return nil
	}

	err := ioutil.WriteFile(zipPath, terminalnotifier(), 0700)
	if err != nil {
		return errors.New("could not write terminal-notifier file")
	}

	defer os.Remove(zipPath)

	err = unpackZipArchive(zipPath, rootPath)
	if err != nil {
		return errors.New("could not unpack zip terminal-notifier file")
	}

	err = os.Chmod(filepath.Join(rootPath, executablePath), 0755)
	if err != nil {
		return errors.New("could not make terminal-notfier executable")
	}

	return nil
}

func unpackZipArchive(filename, tempPath string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}

	defer reader.Close()

	for _, zipFile := range reader.Reader.File {
		name := zipFile.Name
		mode := zipFile.Mode()
		if mode.IsDir() {
			if err = os.MkdirAll(filepath.Join(tempPath, name), 0755); err != nil {
				return err
			}
		} else {
			if err = unpackZippedFile(name, tempPath, zipFile); err != nil {
				return err
			}
		}
	}

	return nil
}

func unpackZippedFile(filename, tempPath string, zipFile *zip.File) error {
	writer, err := os.Create(filepath.Join(tempPath, filename))

	if err != nil {
		return err
	}

	defer writer.Close()

	reader, err := zipFile.Open()
	if err != nil {
		return err
	}

	defer reader.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}

	return nil
}
