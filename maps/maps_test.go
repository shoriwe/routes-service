package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMap() *Map {
	m := New()
	m.Nodes = append(m.Nodes,
		&Node{
			Latitude:  1,
			Longitude: 1,
		},
		&Node{
			Latitude:  4,
			Longitude: 2,
		},
		&Node{
			Latitude:  3,
			Longitude: 4,
		},
		&Node{
			Latitude:  0,
			Longitude: 0,
		},
		&Node{
			Latitude:  2,
			Longitude: 4,
		},
		&Node{
			Latitude:  1,
			Longitude: 1,
		},
	)
	m.Edges = append(m.Edges,
		&Edge{
			Name:       "street 1",
			StartNode:  0,
			EndNode:    1,
			Length:     50,
			SpeedLimit: 50,
		},
		&Edge{
			Name:       "street 1",
			StartNode:  1,
			EndNode:    0,
			Length:     50,
			SpeedLimit: 50,
		},
		&Edge{
			Name:       "street 2",
			StartNode:  0,
			EndNode:    2,
			Length:     20,
			SpeedLimit: 50,
		},
		&Edge{
			Name:       "street 3",
			StartNode:  2,
			EndNode:    1,
			Length:     10,
			SpeedLimit: 10,
		},
		&Edge{
			Name:       "street 4",
			StartNode:  1,
			EndNode:    3,
			Length:     10,
			SpeedLimit: 10,
		},
		&Edge{
			Name:       "street 5",
			StartNode:  3,
			EndNode:    1,
			Length:     10,
			SpeedLimit: 10,
		},
		&Edge{
			Name:       "street 6",
			StartNode:  3,
			EndNode:    4,
			Length:     40,
			SpeedLimit: 30,
		},
		&Edge{
			Name:       "street 7",
			StartNode:  4,
			EndNode:    3,
			Length:     40,
			SpeedLimit: 30,
		},
		&Edge{
			Name:       "street 8",
			StartNode:  4,
			EndNode:    5,
			Length:     2,
			SpeedLimit: 10,
		},
		&Edge{
			Name:       "street 9",
			StartNode:  5,
			EndNode:    4,
			Length:     2,
			SpeedLimit: 10,
		},
		&Edge{
			Name:       "street 10",
			StartNode:  5,
			EndNode:    3,
			Length:     50,
			SpeedLimit: 20,
		},
	)
	return m
}

func TestMap(t *testing.T) {
	testMap()
}

func TestMap_AddEdge(t *testing.T) {
	m := testMap()
	m.AddEdge(
		&Edge{
			Name:       "street 11",
			StartNode:  3,
			EndNode:    5,
			Length:     50,
			SpeedLimit: 20,
		},
	)
	edge := m.Edges[len(m.Edges)-1]
	assert.Equal(t, "street 11", edge.Name)
}

func TestMap_Dijkstra(t *testing.T) {
	m := testMap()
	t.Run("From 0 to 5", func(tt *testing.T) {
		route, distance := m.Dijkstra(0, 5)
		assert.NotZero(tt, distance)
		assert.Equal(tt, []int{0, 2, 1, 3, 4, 5}, route)
	})
	t.Run("From 5 to 2", func(tt *testing.T) {
		route, distance := m.Dijkstra(5, 2)
		assert.NotZero(tt, distance)
		assert.Equal(tt, []int{5, 4, 3, 1, 0, 2}, route)
	})
	t.Run("From 2 to 0", func(tt *testing.T) {
		route, distance := m.Dijkstra(2, 0)
		assert.NotZero(tt, distance)
		assert.Equal(tt, []int{2, 1, 0}, route)
	})
}
