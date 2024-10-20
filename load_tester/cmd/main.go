package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	writer2 "tbl/writer"
)

var address = "127.0.0.1:4000"

func main() {
	logger := initLogger()
	for i := 0; i < 1; i++ {
		logger.Error(fmt.Sprintf("testMessage %v", i))
	}
}

func initLogger() *slog.Logger {
	level := slog.LevelError
	options := &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
	}

	//writer := io.MultiWriter(os.Stdout)

	//w, _ := gelf.NewWriter(address)
	w, err := writer2.NewTblWriter(address)

	if err != nil {
		//return err
		panic(err.Error())
	}

	log.SetOutput(w)

	writer := io.MultiWriter(os.Stdout, w)
	//writer = io.MultiWriter(w)

	logger := slog.New(slog.NewJSONHandler(writer, options))
	//logger := slog.New(slog.NewTextHandler(writer, options))

	logger = logger.With("stream", fmt.Sprintf("loadTester-%s", "1"))

	return logger
}
