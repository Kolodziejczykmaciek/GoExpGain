package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type hashFunc func (string) string  

type Machine struct {
	nameHashFunc hashFunc
}

func (m *Machine) runJob(name string) error {
	newhash := m.nameHashFunc(name)
	fmt.Printf("String \"%s\" (sha256): %s", name, newhash)
	
	return nil
}

func hashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	hashToString := hex.EncodeToString(hash[:])

	return hashToString
}

func main() {
	fmt.Println("Hello")
	m := &Machine{
		nameHashFunc: hashString,
	}
	m.runJob("Ignacy")
}