package client

import (
	"bufio"
	"context"
	"fmt"
	"handleVM/handledb"
	"log"
	"os"
	"strconv"
)

func printOptions() {
	fmt.Println("Enter a number to do an action: \n 1 - Add a new virtual machine \n 2 - Get a virtual machine by ID \n 3 - Exit client")
}

func clientAddVirtualMachine(db *handledb.MySqlStore, scanner *bufio.Scanner) {
	fmt.Println("Adding new machine...")
	fmt.Println("What is the name of the owner?")
	scanner.Scan()
	ownerName := scanner.Text()
	fmt.Println("What operating system do you require?")
	scanner.Scan()
	os := scanner.Text()
	m := handledb.Machine{OperatingSystem: os}
	m.Owner = ownerName
	db.AddVirtualMachine(m)
	fmt.Println("-------------------------------------")
	fmt.Printf("Added Virtual Machine with Owner: %s, Operating System: %s \n", m.Owner, m.OperatingSystem)
	fmt.Println("-------------------------------------")
	fmt.Println("Press enter for more options.")
	scanner.Scan()
	printOptions()
}

func clientGetVirtualMachineById(ctx context.Context, db *handledb.MySqlStore, scanner *bufio.Scanner) {
	fmt.Println("Getting a virtual machine by ID")
	fmt.Println("What is the id of the virtual machine?")

	scanner.Scan()
	id := scanner.Text()
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	res, err := db.GetVirtualMachineById(ctx, intId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-------------------------------------")
	fmt.Println("Accessed the virtual machine: ", res)
	fmt.Println("-------------------------------------")
	fmt.Println("Press enter for more options.")
	scanner.Scan()
	printOptions()
}

func RunClient(ctx context.Context, db *handledb.MySqlStore) {
	fmt.Println("-------------------------------------")
	fmt.Println("Welcome to Alan's VM creation client!")
	fmt.Println("-------------------------------------")
	printOptions()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-ctx.Done():

			return
		default:
			switch scanner.Text() {
			case "1":
				clientAddVirtualMachine(db, scanner)
			case "2":
				clientGetVirtualMachineById(ctx, db, scanner)
			case "3":
				fmt.Println("Goodbye!")
				return
			default:
				fmt.Println("This is not an option.")
				fmt.Println("Press enter for more options.")
			}
		}
	}

	if scanner.Err() != nil {
		// Handle error.
		log.Fatal("Error in user input")
	}

}
