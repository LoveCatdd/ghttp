package ghttp

import "strings"

// route node struct
type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

// 匹配第一个路由节点 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part && child.isWild {
			return child
		}
	}
	return nil
}

// 匹配所有路由节点 用于查找
func (n *node) matchChildren(part string) (nodes []*node) {
	for _, child := range n.children {
		if child.part == part && child.isWild {
			nodes = append(nodes, child)
		}
	}
	return
}

// 插入
func (n *node) insert(pattern string, parts []string, idx int) {
	if len(parts) == idx {
		n.pattern = pattern
		return
	}

	part := parts[idx]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, idx+1)
}

// 查找
func (n *node) search(parts []string, idx int) *node {
	if len(parts) == idx || strings.HasSuffix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[idx]
	children := n.matchChildren(part)

	for _, child := range children {
		res := child.search(parts, idx+1)
		if res != nil {
			return res
		}
	}
	return nil
}
