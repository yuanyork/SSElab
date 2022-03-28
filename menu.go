package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	CMD_MAX_LEN = 128
	DESC_LEN    = 1024
	CMD_NUM     = 10
)

type DataNode struct {
	cmd     string
	desc    string
	handler func() int
	next    *DataNode
}

var head []DataNode = []DataNode{
	DataNode{"help", "this is help cmd!", help, nil},
	DataNode{"version", "menu program v1.0", nil, nil},
	DataNode{"quit", "Quit from menu", Quit, nil},
}

func main() {
	/* cmd line begins */
	head[0].handler = help
	head[0].next = &head[1]
	head[1].next = &head[2]
	for {
		var cmd string
		fmt.Printf("Input a cmd name > ")
		fmt.Scanf("%s", &cmd)
		var p *DataNode = &head[0]
		for p != nil {

			if strings.Compare(p.cmd, cmd) == 0 {
				fmt.Printf("%s - %s\n", p.cmd, p.desc)
				if p.handler != nil {
					p.handler()
				}
				break
			}
			p = p.next
		}
		if p == nil {
			fmt.Printf("This is a wrong cmd!\n ")
		}
	}
}

func help() int {
	fmt.Printf("Menu List:\n")
	var p *DataNode = (*DataNode)head
	for p != nil {
		fmt.Printf("%s - %s\n", p.cmd, p.desc)
		p = p.next
	}
	return 0
}

func Quit() int {
	os.Exit(0)
	return 0
}
