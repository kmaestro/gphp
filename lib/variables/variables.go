package valiables

import (
	"math"
	"sync"
)

var (
	variables map[string]float32
	once      sync.Once
)

func initializeVariables() {
	variables = map[string]float32{
		"PI": math.Pi,
	}
}

func IsExists(key string) bool {
	once.Do(initializeVariables)
	_, exists := variables[key]
	return exists
}

func Get(key string) float32 {
	once.Do(initializeVariables)
	if val, exists := variables[key]; exists {
		return val
	}
	return 0
}

func Set(key string, value float32) {
	once.Do(initializeVariables)
	variables[key] = value
}
