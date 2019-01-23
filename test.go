package main

import (
  "github.com/msaun008/data-structures/binaryTree"
)

func main() {
  tree := binaryTree.NewTree()
  tree.Print()
  tree.Add(3)
  tree.Add(1)
  tree.Add(2)
  tree.Print()
}
