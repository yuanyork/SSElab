package main

import (
	"sync"
)

type LinkTableNode struct {
	pNext *LinkTableNode
}

type LinkTable struct {
	pHead     *LinkTableNode
	pTail     *LinkTableNode
	SumOfNode int
	mutex     *sync.Mutex
}

type contrl interface {
	CreateLinktable() *LinkTable
	DeleteLinktable() bool
	AddLinktableNode(*LinkTableNode) bool
	DeleteLinktableNode(*LinkTableNode) bool
	GetLinktableHead() *LinkTableNode
	GetNextLinktableNode(*LinkTableNode) bool
}

func (l *LinkTable) AddLinktableNode(lNode *LinkTableNode) bool {
	if lNode == nil {
		return false
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.SumOfNode++
	if l.pHead == nil {
		l.pHead = lNode
	}
	if l.pTail == nil {
		l.pTail = lNode
	} else {
		l.pTail.pNext = lNode
		l.pTail = lNode
	}
	return true
}

func (l *LinkTable) DeleteLinktableNode(lNode *LinkTableNode) bool {
	if lNode == nil || l.pHead == nil {
		return false
	}
	l.mutex.Lock()
	defer func() {
		if l.SumOfNode == 0 {
			l.pTail = nil
		}
		l.mutex.Unlock()
	}()
	tmp := l.pHead
	l.SumOfNode--
	if lNode == l.pHead {
		l.pHead = l.pHead.pNext
		return true
	}
	for tmp.pNext != nil {

		if tmp.pNext == lNode {
			if lNode == l.pTail {
				l.pTail = tmp
				l.pTail = nil
				return true
			}
			tmp.pNext = tmp.pNext.pNext
			return true
		}
		tmp = tmp.pNext
	}
	return false
}

func (l *LinkTable) GetLinktableHead() *LinkTableNode {
	return l.pHead
}

func (l *LinkTable) GetNextLinktableNode(pNode *LinkTableNode) *LinkTableNode {
	if pNode == nil {
		return nil
	}
	pTemp := l.pHead
	for pTemp != nil {
		if pTemp == pNode {
			return pTemp.pNext
		}
		pTemp = pTemp.pNext
	}
	return nil
}

func (l *LinkTable) SearchLinkeTableNode(condition func(*LinkTableNode) bool) *LinkTableNode {
	var pNode *LinkTableNode
	pNode = l.pHead
	if pNode == nil || condition == nil {
		return nil
	}
	for pNode != nil {
		if condition(pNode) == true {
			return pNode
		}
		pNode = pNode.pNext
	}
	return nil
}
