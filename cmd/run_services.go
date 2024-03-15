package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	services := []string{}

	projectDir := os.Getenv("PROJECT_DIR")
	dirs, _ := os.ReadDir(fmt.Sprintf("%s/services", projectDir))
	for _, dir := range dirs {
		if dir.IsDir() {
			services = append(services, dir.Name())
		}
	}

	var quitSignal = make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt, syscall.SIGTERM)

	var cmds []*exec.Cmd
	defer func() {
		for _, c := range cmds {
			c.Wait()
			c.Process.Kill()
		}
	}()

	for _, svc := range services {
		log.Println("Running service: ", svc)

		cmd := exec.Command(
			"/bin/sh",
			"-c",
			fmt.Sprintf("cd %s;go run %s/services/%s/cmd/main.go", projectDir, projectDir, svc),
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmds = append(cmds, cmd)

		if err := cmd.Start(); err != nil {
			log.Fatal("Running service error: %v", err)
		}
	}

	<-quitSignal
}
