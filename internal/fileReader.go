package internal

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func Readfile(file string) (string, error) {
	fileContents, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist")
		}
		return "", err
	}
	defer fileContents.Close()

	content, err := io.ReadAll(fileContents)
	if err != nil {
		return "", err
	}
	slog.Info("File", "Content", content)

	return string(content), nil
}
