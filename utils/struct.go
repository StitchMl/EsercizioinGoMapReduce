package utils

// Node Define the structure to represent a node
type Node struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

type Type struct {
	Master  []Node `json:"master"`
	Mapper  []Node `json:"mapper"`
	Reducer []Node `json:"reducer"`
}

// Config Define configuration structure
type Config struct {
	Nodes []Type `json:"nodes"`
}

// MapReduceRequest Request
type MapReduceRequest struct {
	Id int
}

// MapReduceReply Reply
type MapReduceReply struct {
	Chunk []int
}
