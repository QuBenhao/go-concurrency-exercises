//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Create a process
	proc := MockProcess{}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt)
	stopCount := 0

	// Run the process (blocking)
	go proc.Run()
	for {
		select {
		case <-c:
			stopCount++
			if stopCount == 1 {
				fmt.Printf("\nReceived SIGINT, trying to stop the process gracefully")
				go proc.Stop()
			} else {
				fmt.Println("\nReceived second SIGINT, exiting immediately")
				os.Exit(1) // Last resort, kill the program
			}
		}
	}
}
