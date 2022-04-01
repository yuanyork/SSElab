package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"unsafe"
)

const (
	CMD_MAX_LEN = 128
	DESC_LEN    = 1024
	CMD_NUM     = 10
)

type DataNode struct {
	next    LinkTableNode
	cmd     string
	desc    string
	handler func() int
}

type pDataNode *DataNode

var pLinktbl *LinkTable

func InitMenuData() {
	pNode := &DataNode{
		cmd:     "help",
		desc:    "this is help cmd!",
		handler: help,
	}
	pLinktbl.AddLinktableNode((*LinkTableNode)(unsafe.Pointer(pNode)))
	pNode = &DataNode{
		cmd:     "version",
		desc:    "menu program v1.0",
		handler: nil,
	}
	pLinktbl.AddLinktableNode((*LinkTableNode)(unsafe.Pointer(pNode)))
	pNode = &DataNode{
		cmd:     "quit",
		desc:    "Quit from menu",
		handler: Quit,
	}
	pLinktbl.AddLinktableNode((*LinkTableNode)(unsafe.Pointer(pNode)))
}

// var head []DataNode = []DataNode{
// 	DataNode{"help", "this is help cmd!", help, nil},
// 	DataNode{"version", "menu program v1.0", nil, nil},
// 	DataNode{"quit", "Quit from menu", Quit, nil},
// }

func FindCmd(cmd string) pDataNode {
	p := (pDataNode)(unsafe.Pointer(pLinktbl.GetLinktableHead()))
	if p == nil {
		return nil
	}
	for p != nil {
		if strings.Compare(p.cmd, cmd) == 0 {
			return p
		}
		p = (pDataNode)(unsafe.Pointer(pLinktbl.GetNextLinktableNode((*LinkTableNode)(unsafe.Pointer(p)))))
	}
	return nil
}

func main() {
	/* cmd line begins */
	pLinktbl = &LinkTable{
		pHead:     nil,
		pTail:     nil,
		SumOfNode: 0,
		mutex:     &sync.Mutex{},
	}
	InitMenuData()
	for {
		var cmd string
		for {
			fmt.Print("Input a cmd name > ")
			fmt.Scanf("%s", &cmd)
			p := FindCmd(cmd)
			if p == nil {
				fmt.Printf("This is a wrong cmd!\n")
				continue
			}
			fmt.Printf("%s - %s\n", p.cmd, p.desc)
			if p.handler != nil {
				p.handler()
			}
		}
	}
}

func help() int {
	fmt.Printf("Menu List:\n")
	var p *DataNode = (*DataNode)(unsafe.Pointer(pLinktbl.GetLinktableHead()))
	for p != nil {
		fmt.Printf("%s - %s\n", p.cmd, p.desc)
		p = (*DataNode)(unsafe.Pointer(pLinktbl.GetNextLinktableNode((*LinkTableNode)(unsafe.Pointer(p)))))
	}
	return 0
}

func Quit() int {
	os.Exit(0)
	return 0
}
