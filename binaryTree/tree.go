package binaryTree

import (
  "fmt"
)

type Tree struct {
  Root *Node
  Count int
}

func NewTree() *Tree {
  return &Tree{Root: nil, Count: 0}
}

func (t *Tree) Add(data int) bool {
  var node *Node = NewNode(data)
  var prev, current *Node = nil, t.Root

  for current != nil {
    prev = current
    if node.Data < current.Data {
      current = current.Left
    } else {
      current = current.Right
    }
  }

  node.Parent = prev
  if prev == nil {
    t.Root = node
  } else if node.Data < prev.Data {
    prev.Left = node
  } else {
    prev.Right = node
  }

  t.Count++
  return true
}

func (t *Tree) Print() {
  fmt.Printf("Tree: %d\n", t.Count)
  printInOrder(t.Root)
  fmt.Println("")
}

func printInOrder(n *Node) {
  if n == nil { return }

  printInOrder(n.Left)
  fmt.Printf("%d ", n.Data)
  printInOrder(n.Right)
}
