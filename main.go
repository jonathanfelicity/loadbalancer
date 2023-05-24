package main

import (
	"fmt"
	"math/rand"
)

type LoadBalancer struct {
	servers       []string
	serverWeights map[string]int
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		servers:       make([]string, 0),
		serverWeights: make(map[string]int),
	}
}

func (lb *LoadBalancer) AddServer(server string, weight int) {
	lb.servers = append(lb.servers, server)
	lb.serverWeights[server] = weight
}

func (lb *LoadBalancer) RemoveServer(server string) {
	if index, ok := lb.findServerIndex(server); ok {
		lb.servers = append(lb.servers[:index], lb.servers[index+1:]...)
		delete(lb.serverWeights, server)
	}
}

func (lb *LoadBalancer) GetServer() string {
	totalWeight := 0
	for _, weight := range lb.serverWeights {
		totalWeight += weight
	}

	randomNum := rand.Intn(totalWeight)
	currentWeight := 0

	for server, weight := range lb.serverWeights {
		currentWeight += weight
		if randomNum < currentWeight {
			return server
		}
	}

	return ""
}

func (lb *LoadBalancer) findServerIndex(server string) (int, bool) {
	for i, s := range lb.servers {
		if s == server {
			return i, true
		}
	}
	return -1, false
}

func main() {
	lb := NewLoadBalancer()
	lb.AddServer("Server A", 3)
	lb.AddServer("Server B", 2)
	lb.AddServer("Server C", 1)

	for i := 0; i < 10; i++ {
		server := lb.GetServer()
		fmt.Println("Request sent to:", server)
	}
}
