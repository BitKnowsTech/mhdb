//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/magefile/mage/mg"
)

func BuildDB() error {
	fmt.Println("Test DB install...")
	return nil
}

func BuildBin() error {
	fmt.Println("Building API binary...")

	appName := "mhdb"
	if runtime.GOOS == "windows" {
		appName += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", fmt.Sprintf("./build/%s", appName), "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func BuildFull() error {
	mg.Deps(BuildDB)
	mg.Deps(BuildBin)

	return nil
}

func Run() {
	cmd := exec.Command("./build/mhdb")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
