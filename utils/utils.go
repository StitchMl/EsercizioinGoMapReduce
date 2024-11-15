package utils

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func CheckError(info string, err error) error {
	// If an error is returned, print it to the console
	// and exit
	if err != nil {
		log.Fatalf(info+": %v", err)
		return err
	}
	return nil
}

func PrintState(info string) {
	// Print the status of the Cyan colour programme
	fmt.Printf(FgCyan + info + Reset)
}

func PrintWarning(warning string) {
	// Print the status of the Cyan colour programme
	fmt.Printf(FgYellow + warning + Reset)
}

func PrintSuccess(warning string) {
	// Print the status of the Cyan colour programme
	fmt.Printf(FgGreen + warning + Reset)
}

// StartServer starts the server with the information in Node
func StartServer(config Node, handler any) error {
	// Create a new RPC server
	server := rpc.NewServer()

	// Register the handler with the server
	err := server.Register(handler)
	if CheckError("Error registering handler", err) != nil {
		return err
	}

	// Define the address and port for the server
	addr := config.IP + ":" + config.Port

	// Create a listener for incoming connections
	lis, err := net.Listen("tcp", addr)
	if CheckError("Error listening on "+addr, err) != nil {
		return err
	}
	PrintState("RPC server listening...\n")
	fmt.Printf("Starting server at %s:%s\n", config.IP, config.Port)

	// Goroutine to handle incoming connections
	go func() {
		for {
			// Accept a new connection and handle it
			conn, err := lis.Accept()
			_ = CheckError("Error accepting connection", err)
			PrintSuccess("Server listening on " + lis.Addr().String() + "\n")
			PrintWarning("Accepting connection from " + conn.RemoteAddr().String() + "\n")

			// Serve the connection using the RPC server
			go server.ServeConn(conn)
		}
	}()

	// Block forever to keep the server running
	select {}
}

func StartClient(config Node) (*rpc.Client, error) {
	// RPC server address and port
	addr := config.IP + ":" + config.Port
	client, err := rpc.Dial("tcp", addr) // client creation
	if err != nil {
		fmt.Printf("Error connecting to RPC server: %v\n", err)
		return nil, err
	}
	defer func() {
		if client != nil {
			_ = CheckError("Error connecting to RPC server", client.Close())
		}
	}()

	PrintWarning("RPC client listening on " + addr + "\n")

	return client, nil
}
