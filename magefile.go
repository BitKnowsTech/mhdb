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
	cmd := exec.Command("go", "run", "./cmd/database/build")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func BuildBin() error {
	fmt.Println("Building API binary...")

	appName := "mhdb"
	if runtime.GOOS == "windows" {
		appName += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", fmt.Sprintf("./build/%s", appName), ".")
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
	os.Chdir("./build")
	cmd := exec.Command("./mhdb")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
