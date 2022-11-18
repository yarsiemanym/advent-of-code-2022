package common

import (
	"io"
	"net/http"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

func ReadFile(path string) string {
	log.Debug("Reading file.")
	log.Tracef("path = \"%v\"", path)
	bytes, err := os.ReadFile(path)
	Check(err)
	log.Debugf("%v bytes read from file.", len(bytes))
	log.Debug("Converting bytes to text.")
	text := string(bytes)
	log.Debugf("Text is %v characters long.", len(text))
	log.Tracef("Text is \"%v\".", Peek(text, PEEK_MAX_DEFAULT))
	return text
}

func DownloadFile(source string, target string) {
	log.Debug("Downloading file.")
	log.Tracef("source = \"%v\"", source)
	log.Tracef("target = \"%v\"", target)

	if len(Session) == 0 {
		log.Panic("Cannot download puzzle input because the AOC_SESSION_TOKEN environment variable is not set.")
	}

	file, err := os.Create(target)

	if os.IsNotExist(err) {
		dir := path.Dir(target)
		os.Mkdir(dir, 0766)
		file, err = os.Create(target)
	}

	Check(err)

	client := http.Client{}

	cookie := &http.Cookie{
		Name:     "session",
		Value:    Session,
		Domain:   ".adventofcode.com",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	req, err := http.NewRequest("GET", source, nil)
	Check(err)

	req.AddCookie(cookie)

	resp, err := client.Do(req)
	Check(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	Check(err)

	defer file.Close()

	log.Debugf("%v bytes downloaded.", size)
	log.Debug("Input file saved.")
}
