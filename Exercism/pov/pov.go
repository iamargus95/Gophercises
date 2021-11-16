package pov

type Node struct {
	name     string
	children map[string]*Node
}

type Graph struct {
	root  *Node
	nodes map[string]*Node
}

func New() *Graph {
	var res Graph
	res.nodes = make(map[string]*Node)
	return &res
}

func (g *Graph) AddNode(nodeLabel string) {
	g.nodes[nodeLabel] = &Node{nodeLabel, make(map[string]*Node)}
}

func (g *Graph) contains(name string) bool {
	_, exists := g.nodes[name]
	return exists
}

func (g *Graph) AddArc(from, to string) {
	if !g.contains(from) {
		g.AddNode(from)
	}
	src := g.nodes[from]
	dst := g.nodes[to]
	src.children[dst.name] = dst
	if g.root == nil {
		g.root = src
	}
}

func (g *Graph) ArcList() []string {
	arcs := make([]string, 0, len(g.nodes))
	for _, node := range g.nodes {
		for _, child := range node.children {
			arcStr := node.name + " -> " + child.name
			arcs = append(arcs, arcStr)
		}
	}
	return arcs
}

// find path to dest; depth first search
func findPath(current, dest *Node, log *map[string]bool) (bool, []*Node) {
	if current.name == dest.name {
		return true, []*Node{current}
	}
	for _, child := range current.children {
		if _, visited := (*log)[child.name]; visited {
			continue
		} else {
			(*log)[child.name] = true
		}
		found, subpath := findPath(child, dest, log)
		if found {
			path := []*Node{current}
			path = append(path, subpath...)
			return true, path
		}
	}
	return false, []*Node{}
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	visited := make(map[string]bool)
	_, path := findPath(g.nodes[oldRoot], g.nodes[newRoot], &visited)
	for i, node := range path {
		if i > 0 {
			prev := path[i-1]
			// add previous to children
			node.children[prev.name] = prev
			// remove from previous's children
			delete(prev.children, node.name)
		}
	}
	return g
}
