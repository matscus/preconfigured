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
	for _, v := range dirs {
		if v.IsDir() {
			if v.Name() != "vendor" && v.Name() != path && v.Name() != ".git" {
				err = filepath.WalkDir(v.Name(),
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
							readFile, err := ioutil.ReadFile(filePath)
							if err != nil {
								return err
							}
							data := []byte(os.ExpandEnv(string(readFile)))
							writeFile, err := os.Create(fmt.Sprintf("%s/%s", path, filePath))
							if err != nil {
								return err
							}
							writeFile.Write(data)
							writeFile.Close()
						}
						return nil
					})
				if err != nil {
					log.Error(err)
				}
			}
		} else {
			err := os.MkdirAll(fmt.Sprintf("%s/%s", path, service), 0777)
			if err != nil {
				log.Error(err)
			}
			readFile, err := ioutil.ReadFile(service + "/" + v.Name())
			if err != nil {
				log.Error(err)
			}
			data := []byte(os.ExpandEnv(string(readFile)))
			writeFile, err := os.Create(fmt.Sprintf("%s/%s/%s", path, service, v.Name()))
			if err != nil {
				log.Error(err)
			}
			writeFile.Write(data)
			writeFile.Close()
		}
	}
	if err != nil {
		log.Fatal(err)
	}
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
