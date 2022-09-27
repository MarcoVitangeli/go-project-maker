package cli

import (
	"fmt"
	"os"
	"os/exec"
)

type CommandRunner struct {
}

const (
	MainFileName    = "main.go"
	MainFileContent = "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello World!\")\n}\n\n"
)

func NewCommandRunner() *CommandRunner {
	return &CommandRunner{}
}

func (c *CommandRunner) GoTo(dirName string) error {
	err := os.Chdir(dirName)

	if err != nil {
		return err
	}

	return nil
}

func (c *CommandRunner) GoHome() error {
	home, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	return c.GoTo(home)
}

func (c *CommandRunner) CreateDir(name string, perm os.FileMode) error {
	err := os.Mkdir(name, perm)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandRunner) CreateDirAndGo(name string, perm os.FileMode) error {
	err := c.CreateDir(name, perm)
	if err != nil {
		return err
	}

	err = c.GoTo(name)

	if err != nil {
		return err
	}

	return nil
}

func (c *CommandRunner) getGitUser() ([]byte, error) {
	cmd := exec.Command("git", "config", "user.name")
	b, err := cmd.Output()

	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *CommandRunner) InitGoMod(name string) error {
	gitUser, err := c.getGitUser()
	gitUser = gitUser[:len(gitUser)-1]
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", fmt.Sprintf("github.com/%s/%s", string(gitUser), name))
	_, err = cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

func (c *CommandRunner) CreateFile(fileName, strToWrite string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	_, err = file.WriteString(strToWrite)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandRunner) CreateMainGoFile() error {
	return c.CreateFile(MainFileName, MainFileContent)
}
