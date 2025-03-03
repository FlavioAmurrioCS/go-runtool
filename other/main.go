package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func extract_tar_gz(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()
	uncompressed_stream, _ := gzip.NewReader(file)
	defer uncompressed_stream.Close()
	tarReader := tar.NewReader(uncompressed_stream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(header.Name, os.FileMode(header.Mode))
		case tar.TypeReg:
			outFile, err := os.OpenFile(header.Name, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			defer outFile.Close()
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				panic(err)
			}
		default:
			panic("unknown type")
		}
	}

}

func main() {
	client := getGithubClient()
	githubSource := NewGithubSourceFromUrl("https://github.com/junegunn/fzf/releases/tag/v0.60.0")
	links := githubSource.links()
	link := bestLink(links)
	req, _ := http.NewRequest("GET", link, nil)
	file := download(client, req, "")

	extract_tar_gz(file)

	command := exec.Command("./fzf")

	command.Run()

}
