package logging

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

var path string = "logs"

func Location(location string) {
	path = location
}

func New() *slog.Logger {

	today := time.Now()
	log_name := fmt.Sprintf("%s/%d-%d-%d.txt", path, today.Year(),today.Month(),today.Day())

	file, err := os.Open(log_name)

	if err != nil {

		file, _ := os.Create(log_name)
		os.Stdout = file

	}

	os.Stdout = file
	return slog.New(slog.NewTextHandler(file, nil))

}