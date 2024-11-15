package mapper

import (
	"esercizio_go_mapreduce/utils"
	"fmt"
	"net/rpc"
)

// Function for a synchronous RPC call
func makeRequestSync(client *rpc.Client, id int) ([]int, error) {
	args := utils.MapReduceRequest{Id: id} // Create RPC Request
	reply := utils.MapReduceReply{}        // Create the structure for the response

	// Synchronous RPC call: pass &reply to get response from server
	err := client.Call("Handler.Compute", args, &reply)
	if utils.CheckError("Error when requesting chunk", err) != nil {
		return nil, err
	}

	return reply.Chunk, nil // Returns the chunk from the response
}

// Function for an asynchronous RPC call
func makeRequestAsync(client *rpc.Client, args interface{}, reply interface{}) *rpc.Call {
	// Asynchronous call
	return client.Go("Handler.Method", args, reply, nil)
}

func Mapper(config *utils.Config, mapperConfig utils.Node, id int) error {
	fmt.Printf("Start %v\n", mapperConfig.Name)
	temp := config.Nodes[0]
	client, err := utils.StartClient(temp.Master[0])
	if utils.CheckError("Error starting server", err) != nil {
		return err
	}

	// Request Parameters
	args := id

	// Synchronous Calling
	chunk, err := makeRequestSync(client, args)
	if utils.CheckError("Error in synchronous RPC call", err) != nil {
		return err
	}
	utils.PrintSuccess("Result (Sync): " + fmt.Sprint(chunk) + "\n")

	/*// Asynchronous call
	asyncCall := makeRequestAsync(client, args, &reply)
	<-asyncCall.Done
	if CheckError("Error in asynchronous RPC call", asyncCall.Error) != nil {
		return asyncCall.Error
	}
	fmt.Printf("Result (Async): %v\n", reply)*/

	return nil
}
