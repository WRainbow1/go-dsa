package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	graph_copy := Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	visited[node] = &graph_copy

	graph_copy.Neighbors = create_node_children(node.Neighbors, visited)

	return &graph_copy
}

func create_node_children(nodes []*Node, visited map[*Node]*Node) []*Node {

	var output []*Node

	for _, node := range nodes {

		if copy, exists := visited[node]; exists {
			output = append(output, copy)
			continue
		}

		node_copy := Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}

		visited[node] = &node_copy
		node_copy.Neighbors = create_node_children(node.Neighbors, visited)
		output = append(output, &node_copy)
	}

	return output
}
