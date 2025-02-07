// Challenge: Walk the Dog
// Task: given a list of actions with random durations implement a function that
//		simulates the concurrent execution of th ordered list of actions.

package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const MAX_SECOND = 3

type Dog struct {
	name string
}

func (d Dog) fetchLeash() {
	log.Printf("%s goes to fetch leash.\n", d.name)
	randomSleep()
	log.Printf("%s has fetched leash. Woof woof!\n", d.name)
}

func (d Dog) findTreats() {
	log.Printf("%s goes to fetch treats.\n", d.name)
	randomSleep()
	log.Printf("%s has fetches the treats. Woof woof!\n", d.name)
}

func (d Dog) runOutside() {
	log.Printf("%s starts running outside.\n", d.name)
	randomSleep()
	log.Printf("%s is having fun outside. Woof woof!\n", d.name)
}

type Owner struct {
	name string
}

func (o Owner) putShoesOn() {
	log.Printf("%s starts putting shoes on.\n", o.name)
	randomSleep()
	log.Printf("%s finishes putting shoes on.\n", o.name)
}

func (o Owner) findKeys() {
	log.Printf("%s starts looking for keys.\n", o.name)
	randomSleep()
	log.Printf("%s has found keys.\n", o.name)
}

func (o Owner) lockDoor() {
	log.Printf("%s starts locking the door.\n", o.name)
	randomSleep()
	log.Printf("%s has locked the door.\n", o.name)
}

func randomSleep() {
	r := rand.Intn(MAX_SECOND)
	time.Sleep(time.Duration(r)*time.Second + 500*time.Millisecond)
}

func executeWalk(ownerActions []func(), dogActions []func()) {
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	executionActions := func(actions []func()) {
		defer wg.Done()
		for _, action := range actions {
			action()
		}
	}

	go executionActions(ownerActions)
	go executionActions(dogActions)
}

func main() {
	owner := Owner{name: "Jimmy"}
	dog := Dog{name: "Lucky"}

	ownerActions := []func(){
		owner.putShoesOn,
		owner.findKeys,
		owner.lockDoor,
	}

	dogActions := []func(){
		dog.fetchLeash,
		dog.findTreats,
		dog.runOutside,
	}

	executeWalk(ownerActions, dogActions)
}
