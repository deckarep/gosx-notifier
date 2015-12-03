package gosxnotifier

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
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

func supportedOS() bool {
	if runtime.GOOS == "darwin" {
		return true
	} else {
		log.Print("OS does not support terminal-notifier")
		return false
	}
}

func init() {
	if supportedOS() {
		err := installTerminalNotifier()
		if err != nil {
			log.Fatalf("Could not install Terminal Notifier to a temp directory: %s", err)
		} else {
			FinalPath = filepath.Join(rootPath, executablePath)
		}
	}
}

func exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func installTerminalNotifier() error {
	rootPath = filepath.Join(os.TempDir(), tempDirSuffix)

	//if terminal-notifier.app already installed no-need to re-install
	if exists(filepath.Join(rootPath, executablePath)) {
		return nil
	}

	err := ioutil.WriteFile(zipPath, terminalnotifier(), 0700)
	if err != nil {
		return fmt.Errorf("could not write terminal-notifier file (%s): %s", zipPath, err)
	}

	defer os.Remove(zipPath)

	err = unpackZipArchive(zipPath, rootPath)
	if err != nil {
		return fmt.Errorf("could not unpack zip terminal-notifier file: %s", err)
	}

	err = os.Chmod(filepath.Join(rootPath, executablePath), 0755)
	if err != nil {
		return fmt.Errorf("could not make terminal-notifier executable: %s", err)
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
