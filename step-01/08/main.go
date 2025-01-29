// Challenge: Spreading Gossip
// Task: given a list of nodes implement a function
//		that ensures all nodes are visited

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const FILE = "./step-01/08/friends.json"

// Friend represents a friend and their connections.
type Friend struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
}

// hearGossip indicates that the friend has heard the gossip.
func (f *Friend) hearGossip() {
	log.Printf("%s has heard the gossip!\n", f.Name)
}

// Friends represents the map of friends and connections.
type Friends struct {
	friends map[string]Friend
}

// getFiend fetches the friend given an id.
func (f *Friends) getFriend(id string) Friend {
	return f.friends[id]
}

// getRandomFriend returns a random friend.
func (f *Friends) getRandomFriend() Friend {
	rand.Seed(time.Now().Unix())
	id := (rand.Intn(len(f.friends)-1) + 1) * 100
	return f.getFriend(fmt.Sprint(id))
}

// spreadGossip ensures that all the friends in the map have heard the news
func spreadGossip(root Friend, friends Friends, visited map[string]struct{}) {
	for _, id := range root.Friends {
		if _, isVisited := visited[id]; !isVisited {
			f := friends.getFriend(id)
			f.hearGossip()
			visited[id] = struct{}{}
			spreadGossip(f, friends, visited)
		}
	}
}

// importData reads the input data from file and create the friends map
func importData() Friends {
	file, err := os.ReadFile(FILE)
	if err != nil {
		log.Fatal(err)
	}

	var data []Friend
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	fm := make(map[string]Friend, len(data))
	for _, d := range data {
		fm[d.ID] = d
	}

	return Friends{
		friends: fm,
	}
}

func main() {
	friends := importData()
	root := friends.getRandomFriend()
	root.hearGossip()
	visited := make(map[string]struct{})
	spreadGossip(root, friends, visited)
}
