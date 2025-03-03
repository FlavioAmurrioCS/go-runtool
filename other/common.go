package main

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func download(client *http.Client, req *http.Request, filename string) string {
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if filename == "" {
		_, params, _ := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
		if filenamee, ok := params["filename"]; ok {
			filename = filenamee
		} else {
			path_split := strings.Split(req.URL.Path, "/")
			filename = path_split[len(path_split)-1]
		}
	}
	filename, _ = filepath.Abs(filename)
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
	return filename
}

func bestLink(links []string) string {
	osSpecificLinks := filter(links, func(s string) bool { return strings.Contains(s, runtime.GOOS) })
	if len(osSpecificLinks) == 0 {
		osSpecificLinks = links
	}
	archSpecificLinks := filter(osSpecificLinks, func(s string) bool { return strings.Contains(s, runtime.GOARCH) })
	if len(archSpecificLinks) == 0 {
		archSpecificLinks = osSpecificLinks
	}
	// negative filters
	f := func(s string) bool { return true }
	switch runtime.GOOS {
	case "linux":
		f = func(s string) bool { return !strings.Contains(s, "windows") && !strings.Contains(s, "darwin") }
	case "darwin":
		f = func(s string) bool { return !strings.Contains(s, "windows") && !strings.Contains(s, "linux") }
	case "windows":
		f = func(s string) bool { return !strings.Contains(s, "linux") && !strings.Contains(s, "darwin") }
	}
	negativeFilteredLinks := filter(archSpecificLinks, f)
	if len(negativeFilteredLinks) == 0 {
		negativeFilteredLinks = archSpecificLinks
	}
	return negativeFilteredLinks[0]
}
