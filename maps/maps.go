package maps

import "math"

type (
	Node struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	Edge struct {
		Name       string  `json:"name"`  // Street name
		StartNode  int     `json:"start"` // Node index
		EndNode    int     `json:"end"`   // Node index
		Length     float64 `json:"length"`
		SpeedLimit float64 `json:"speedLimit"`
	}
	Map struct {
		Nodes []*Node `json:"nodes"`
		Edges []*Edge `json:"edges"`
	}
)

func (m *Map) Dijkstra(startNode, endNode int) ([]int, float64) {
	numNodes := len(m.Nodes)
	dist := make([]float64, numNodes)
	prev := make([]int, numNodes)
	visited := make([]bool, numNodes)

	for i := 0; i < numNodes; i++ {
		dist[i] = math.Inf(1)
		prev[i] = -1
	}

	dist[startNode] = 0

	for {
		u := -1
		smallestDist := math.Inf(1)

		for i := 0; i < numNodes; i++ {
			if !visited[i] && dist[i] < smallestDist {
				u = i
				smallestDist = dist[i]
			}
		}

		if u == -1 || u == endNode {
			break
		}

		visited[u] = true

		for _, edge := range m.Edges {
			if edge.StartNode == u {
				v := edge.EndNode
				alt := dist[u] + edge.Length

				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}

	if prev[endNode] == -1 {
		return nil, math.Inf(1)
	}

	path := make([]int, 0)

	for u := endNode; u != -1; u = prev[u] {
		path = append([]int{u}, path...)
	}

	return path, dist[endNode]
}

func (m *Map) AddEdge(edge *Edge) {
	m.Edges = append(m.Edges, edge)
}

func New() *Map {
	return &Map{}
}
