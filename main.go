package main

import (
	"encoding/json"
	"esercizio_go_mapreduce/mapper"
	"esercizio_go_mapreduce/master"
	"esercizio_go_mapreduce/utils"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

// Start-up functions for the various nodes
func startMaster(wg *sync.WaitGroup, config *utils.Config, masterConfig utils.Node) {
	defer wg.Done()
	utils.PrintState("Starting Master...\n")
	// Logic to start the master with the specific configuration
	// Example dataset and chunk size
	dataset := []int{6, 12, 15, 2, 6, 125, 1, 3, 8}
	chunkSize := 3
	err := master.Master(dataset, chunkSize, config, masterConfig)
	if utils.CheckError("Error when starting the master", err) != nil {
		return
	}
}

func startMapper(wg *sync.WaitGroup, config *utils.Config, mapperID int, mapperConfig utils.Node) {
	defer wg.Done()
	utils.PrintState("Starting Mapper " + strconv.Itoa(mapperID) + "...\n")
	err := mapper.Mapper(config, mapperConfig, mapperID)
	if utils.CheckError("Error when starting the mapper "+strconv.Itoa(mapperID), err) != nil {
		return
	}
}

func startReducer(wg *sync.WaitGroup, config *utils.Config, reducerID int, reducerConfig utils.Node) {
	defer wg.Done()
	fmt.Printf("Starting Reducer %d...\n", reducerID)
	// Logic to start the reducer with specific ID
	// ...
}

// Function for loading configuration from JSON file
func loadConfig(filename string) (*utils.Config, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if utils.CheckError("_", err) != nil {
			return
		}
	}(file)

	// Read the contents of the file
	bytes, err := io.ReadAll(file)
	if utils.CheckError("_", err) != nil {
		return nil, err
	}

	// Decode the JSON in the Config
	var config utils.Config
	if err := json.Unmarshal(bytes, &config); utils.CheckError("_", err) != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	// Upload configuration
	config, err := loadConfig("config.json")
	if utils.CheckError("Error loading configuration", err) != nil {
		return
	}

	// Print the details of each node for verification
	for _, nodes := range config.Nodes {
		for _, node := range nodes.Master {
			fmt.Printf("MASTER -> Nome: %s, IP: %s, Porta: %s\n", node.Name, node.IP, node.Port)
		}
		for _, node := range nodes.Mapper {
			fmt.Printf("MAPPER -> Nome: %s, IP: %s, Porta: %s\n", node.Name, node.IP, node.Port)
		}
		for _, node := range nodes.Reducer {
			fmt.Printf("REDUCER -> Nome: %s, IP: %s, Porta: %s\n", node.Name, node.IP, node.Port)
		}
	}

	var wg sync.WaitGroup
	temp := config.Nodes

	// Start the Master
	wg.Add(1)
	go startMaster(&wg, config, temp[0].Master[0])

	// Start the mappers
	for i, mapperConfig := range temp[1].Mapper {
		wg.Add(1)
		go startMapper(&wg, config, i, mapperConfig)
	}

	// Start the reducers
	for i, reducerConfig := range temp[2].Reducer {
		wg.Add(1)
		go startReducer(&wg, config, i, reducerConfig)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Application completed.")
}
