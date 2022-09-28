package unoconvert

import (
	"fmt"
	"log"
	"os/exec"
)

var unoconvert = &Unoconvert{
	Interface:  "127.0.0.1",
	Port:       "2002",
	Executable: "unoconvert",
}

func SetExecutable(executable string) {
	unoconvert.SetExecutable(executable)
}

func SetInterface(interf string) {
	unoconvert.SetInterface(interf)
}

func SetPort(port string) {
	unoconvert.SetPort(port)
}

func Run(infile string, outfile string, opts ...string) error {
	return unoconvert.Run(infile, outfile, opts...)
}

type Unoconvert struct {
	Interface  string
	Port       string
	Executable string
}

func (u *Unoconvert) SetExecutable(executable string) {
	u.Executable = executable
}

func (u *Unoconvert) SetInterface(interf string) {
	u.Interface = interf
}

func (u *Unoconvert) SetPort(port string) {
	u.Port = port
}

func (u *Unoconvert) Run(infile string, outfile string, opts ...string) error {
	var args = []string{}

	connections := []string{
		fmt.Sprintf("--interface=%s", u.Interface),
		fmt.Sprintf("--port=%s", u.Port),
	}

	files := []string{infile, outfile}

	args = append(connections, files...)
	args = append(args, opts...)

	log.Printf("Command: %s %s", u.Executable, args)
	cmd := exec.Command(u.Executable, args...)

	return cmd.Run()
}
