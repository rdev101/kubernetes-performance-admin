package model

// Node model for api
type Node struct {
	Name string
}

// Nodes (plural) model for api
type Nodes struct {
	Data *[]Nodes
}
