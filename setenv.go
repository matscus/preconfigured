package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	logLevel, service, path string
	isServices              bool
)

func main() {
	flag.StringVar(&service, "service", ".", "PATH service dir")
	flag.StringVar(&path, "path", "result", "PATH from create result files")
	flag.StringVar(&logLevel, "loglevel", "INFO", "log level, default INFO")
	flag.Parse()
	setLogLevel(logLevel)
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file")
	}
	dirs, err := os.ReadDir(service)
	if err != nil {
		log.Fatal(err)
	}
	if service != "." {
		isServices = true
	}
	for _, v := range dirs {
		if v.IsDir() {
			name := v.Name()
			if isServices {
				name = fmt.Sprintf("%s/%s", service, v.Name())
			}
			if v.Name() != "vendor" && v.Name() != path && v.Name() != ".git" {
				err = filepath.WalkDir(name,
					func(filePath string, info os.DirEntry, err error) error {
						if err != nil {
							return err
						}
						if info.IsDir() {
							err = os.MkdirAll(fmt.Sprintf("%s/%s", path, filePath), 0777)
							if err != nil {
								return err
							}
						} else {
							writer(filePath, fmt.Sprintf("%s/%s", path, filePath))
							if err != nil {
								return err
							}
						}
						return nil
					})
				if err != nil {
					log.Error(err)
				}
			}
		} else {
			if isServices {
				err := os.MkdirAll(fmt.Sprintf("%s/%s", path, service), 0777)
				if err != nil {
					log.Error(err)
				}
				writer(fmt.Sprintf("%s/%s", service, v.Name()), fmt.Sprintf("%s/%s/%s", path, service, v.Name()))
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}

func writer(r string, w string) error {
	readFile, err := ioutil.ReadFile(r)
	if err != nil {
		return err
	}
	data := []byte(os.ExpandEnv(string(readFile)))
	writeFile, err := os.Create(w)
	if err != nil {
		return err
	}
	writeFile.Write(data)
	writeFile.Close()
	return nil
}

func setLogLevel(level string) {
	level = strings.ToUpper(level)
	switch level {
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	}
}
