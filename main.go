package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

// CustomNode represents a node with additional attributes
type CustomNode struct {
	IDValue int64  // ID for the node
	Name    string // Custom name
	Prompt  string // Custom YAML attribute
}

// ID returns the node's unique identifier
func (n CustomNode) ID() int64 {
	return n.IDValue
}

// ID returns the node's unique identifier
func (n CustomNode) Edge() int64 {
	return n.IDValue
}

// CustomEdge represents an edge with optional attributes
type CustomEdge struct {
	FromNode graph.Node
	ToNode   graph.Node
}

// From returns the source node of the edge
func (e CustomEdge) From() graph.Node {
	return e.FromNode
}

// To returns the target node of the edge
func (e CustomEdge) To() graph.Node {
	return e.ToNode
}

// To returns the target node of the edge
func (e CustomEdge) ReversedEdge() graph.Edge {
	return nil
}

// DepthFirstSearch performs a DFS starting from the given start node
func DepthFirstSearch(g *simple.DirectedGraph, start graph.Node) {
	visited := make(map[int64]bool) // To keep track of visited nodes
	stack := []graph.Node{start}    // Stack for DFS

	fmt.Println("Depth-First Traversal:")

	for len(stack) > 0 {
		// Pop the last node from the stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node.ID()] {
			continue
		}

		// Mark the node as visited and process it
		visited[node.ID()] = true
		customNode := node.(*CustomNode)
		fmt.Printf("Visited Node ID: %d, Name: %s, Prompt: %s\n", customNode.ID(), customNode.Name, customNode.Prompt)

		// Push all unvisited neighbors onto the stack
		neighbors := g.From(node.ID())
		for neighbors.Next() {
			neighbor := neighbors.Node()
			if !visited[neighbor.ID()] {
				stack = append(stack, neighbor)
			}
		}
	}
}

func main() {
	// Create a directed graph
	g := simple.NewDirectedGraph()

	// Create custom nodes
	node1 := &CustomNode{IDValue: 1, Name: "Start", Prompt: ""}
	node2 := &CustomNode{IDValue: 2, Name: "CreateHaiku", Prompt: "Create a haiku about nature"}
	node3 := &CustomNode{IDValue: 3, Name: "TranslateHaiku", Prompt: "Translate a haiku to pig latin"}

	// Add nodes to the graph
	g.AddNode(node1)
	g.AddNode(node2)
	g.AddNode(node3)

	// Add edges between nodes
	g.SetEdge(CustomEdge{FromNode: node1, ToNode: node2})
	g.SetEdge(CustomEdge{FromNode: node2, ToNode: node3})

	// Traverse and print the graph with attributes
	fmt.Println("Nodes:")
	for nodes := g.Nodes(); nodes.Next(); {
		node := nodes.Node().(*CustomNode)
		fmt.Printf("  ID: %d, Name: %s, Prompt: %s\n", node.ID(), node.Name, node.Prompt)
	}

	fmt.Println("\nEdges:")
	for edges := g.Edges(); edges.Next(); {
		edge := edges.Edge().(CustomEdge)
		fmt.Printf("  From: %d, To: %d\n", edge.From().ID(), edge.To().ID())
	}

	DepthFirstSearch(g, node1)
}
