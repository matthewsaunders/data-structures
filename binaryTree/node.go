package binaryTree

import (
)

type Node struct {
  Data int
  Parent *Node
  Left *Node
  Right *Node
}

func NewNode(data int) *Node {
  return &Node{Data: data, Parent: nil, Left: nil, Right: nil}
}
