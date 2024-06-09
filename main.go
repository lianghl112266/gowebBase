package main

import "goweb/cmd"

func main() {
	cmd.Start()
	defer cmd.Clean()
}
