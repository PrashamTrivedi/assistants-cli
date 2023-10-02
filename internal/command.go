package internal

import (
	"log/slog"
	"os/exec"
	"strings"
)

func RunCommand(command string, args ...string) (string, error) {
	commandToRun := exec.Command(command, args...)
	var out strings.Builder
	commandToRun.Stdout = &out
	err := commandToRun.Run()
	if err != nil {
		slog.Error("Command", "Error", err)
		return "", err
	}
	slog.Info("Command", "Cmd to run", command, "Args", args, "Output", out.String())
	return out.String(), nil
}
