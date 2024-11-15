package master

import (
	"esercizio_go_mapreduce/utils"
	"fmt"
	"strconv"
)

var chunks [][]int

// Handler Generic handler structure (you can change the name to fit your use case)
type Handler struct{}

// Compute A generic method to handle an RPC request (no specific logic here)
func (Handler) Compute(request utils.MapReduceRequest, reply *utils.MapReduceReply) error {
	utils.PrintState("Start sending chunk " + strconv.Itoa(request.Id) + "...\n")
	input := request.Id
	fmt.Printf("Give chunk %d to mapper %d\n", input-1, input)

	// Suppose you have a slice of chunks
	// Check that 'chunks' is defined and populated correctly
	chunk := chunks[input-1]
	reply.Chunk = chunk // Set the chunk in the response

	return nil
}

// PartitionDataset divides a dataset into chunks of uniform size.
// Returns a two-dimensional array in which each sub-array represents a chunk.
func partitionDataset(dataset []int, chunkSize int) [][]int {
	var chunks [][]int

	for i := 0; i < len(dataset); i += chunkSize {
		end := i + chunkSize

		// Check that the final chunk does not exceed the length of the dataset
		if end > len(dataset) {
			end = len(dataset)
		}

		// Adds the chunk to the chunk list
		chunks = append(chunks, dataset[i:end])
	}

	return chunks
}

func Master(dataset []int, chunkSize int, config *utils.Config, masterConfig utils.Node) error {
	// Partition the dataset
	chunks := partitionDataset(dataset, chunkSize)

	// Print chunks for verification
	for i, chunk := range chunks {
		fmt.Printf("Chunk %d: %v\n", i+1, chunk)
	}

	// Instantiate the handler
	handler := new(Handler)

	//fmt.Printf("Starting server master...")
	err := utils.StartServer(masterConfig, *handler)
	if utils.CheckError("Error starting server", err) != nil {
		return err
	}

	return nil
}
