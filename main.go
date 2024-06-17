package main

import (
	"log"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal("ffmpeg not found in PATH")
	}

	if len(os.Args[1:]) < 2 {
		slog.Error("Too few arguments", "len", len(os.Args[1:]))
		log.Fatal("Usage: ffpreset <preset> <input> <input> ...\nAvailable presets: (w)ebp, (a)vif")
	}

	var commands [][]string
	for _, inputFile := range os.Args[2:] {
		outputFile := strings.TrimSuffix(inputFile, path.Ext(inputFile))
		switch strings.ToLower(os.Args[1]) {
		case "w", "webp":
			commands = append(commands, []string{"ffmpeg", "-i", inputFile, "-vcodec", "libwebp", "-quality", "90", outputFile + ".webp"})
		case "a", "avif":
			commands = append(commands, []string{"ffmpeg", "-i", inputFile, "-c:v", "libaom-av1", "-cpu-used", "1", "-crf", "25", outputFile + ".avif"})
		case "j", "jxl":
			if _, err := exec.LookPath("cjxl"); err != nil {
				slog.Error("cjxl not found in PATH", "err", err)
			}
			commands = append(commands, []string{"cjxl", inputFile, outputFile + ".jxl", "-d", "1", "-e", "8"})
		default:
			slog.Error("Unknown preset", "preset", os.Args[1])
			os.Exit(1)
		}
	}

	for _, command := range commands {
		slog.Info("Executing command: ", "command", command)
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
