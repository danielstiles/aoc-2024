package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2024/01/internal/list"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	slog.Info("Answer", slog.Int("total", list.Process2(lines)))
}
