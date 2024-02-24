package wrapper

import (
	"bufio"
	"fmt"
)

type Console interface {
	Start() error
	Kill() error
	WriteCmd(string) error
	ReadLine() (string, error)
}

type DefaultConsole struct {
	Cmd    JavaExec
	Stdout *bufio.Reader
	Stdin  *bufio.Writer
}

func NewConsole(cmd JavaExec) *DefaultConsole {
	c := &DefaultConsole{
		Cmd: cmd,
	}

	c.Stdout = bufio.NewReader(cmd.Stdout())
	c.Stdin = bufio.NewWriter(cmd.Stdin())
	return c
}

func (c *DefaultConsole) Start() error {
	return c.Cmd.Start()
}

func (c *DefaultConsole) Kill() error {
	return c.Cmd.Kill()
}

func (c *DefaultConsole) WriteCmd(cmd string) error {
	wrappedCmd := fmt.Sprintf("%s\r\n", cmd)
	_, err := c.Stdin.WriteString(wrappedCmd)
	if err != nil {
		return err
	}
	return c.Stdin.Flush()
}

func (c *DefaultConsole) ReadLine() (string, error) {
	return c.Stdout.ReadString('\n')
}
