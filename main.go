package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
		case "run":
			run()
		case "child":
			child()
		default:
			fmt.Print("command not found\n")
	}


	
}

func child() {
	fmt.Printf("...input arguments: %d\nrunning %s with PID:%d\n", len(os.Args), os.Args[2], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	must(syscall.Chroot("/home/test"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run())	
}

func run() {
	fmt.Printf("...input arguments: %d\nrunning %s with PID:%d\n", len(os.Args), os.Args[2], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}