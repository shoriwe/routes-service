//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/shoriwe/routes-service/maps"
)

const earthRadius = 6371.0

func radians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := radians(lat2 - lat1)
	dLon := radians(lon2 - lon1)

	lat1 = radians(lat1)
	lat2 = radians(lat2)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLon/2), 2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func GenerateRandomMap(nodesCount, edgesCount int) *maps.Map {
	m := maps.New()
	for i := 0; i < nodesCount; i++ {
		node := &maps.Node{
			Latitude:  rand.Float64()*180 - 90,
			Longitude: rand.Float64()*360 - 180,
		}
		m.Nodes = append(m.Nodes, node)
	}
	for i := 0; i < edgesCount; i++ {
		startIndex := rand.Intn(nodesCount)
		endIndex := rand.Intn(nodesCount)

		for endIndex == startIndex {
			endIndex = rand.Intn(nodesCount)
		}

		edge := &maps.Edge{
			Name:      fmt.Sprintf("Street %d", i+1),
			StartNode: startIndex,
			EndNode:   endIndex,
			Length: haversineDistance(
				m.Nodes[startIndex].Latitude,
				m.Nodes[startIndex].Longitude,
				m.Nodes[endIndex].Latitude,
				m.Nodes[endIndex].Longitude,
			),
			SpeedLimit: rand.Float64()*90 + 30,
		}
		m.AddEdge(edge)
	}

	return m
}

func removeOrphanNodes(m *maps.Map) *maps.Map {
	connectedNodes := make(map[int]bool)

	for _, edge := range m.Edges {
		connectedNodes[edge.StartNode] = true
		connectedNodes[edge.EndNode] = true
	}

	newNodes := make([]*maps.Node, 0)
	for idx := range connectedNodes {
		newNodes = append(newNodes, m.Nodes[idx])
	}

	m.Nodes = newNodes
	return m
}

func main() {
	nodesCount := 300
	edgesCount := nodesCount * 3

	m := GenerateRandomMap(nodesCount, edgesCount)
	m = removeOrphanNodes(m)

	mapJSON, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling map to JSON:", err)
		return
	}

	file, err := os.Create("imaginary_city.json")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(mapJSON)
	if err != nil {
		panic(err)
	}
}
