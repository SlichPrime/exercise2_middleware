package main

import "fmt"

func main() {
	var quota int
	fmt.Print("Set Seminar Quota: ")
	fmt.Scan(&quota)

	type Participant struct{ Name, Status string }
	var list []Participant

	for {
		var name string
		fmt.Print("\nEnter Name (or 'stop'): ")
		fmt.Scan(&name)
		if name == "stop" {
			break
		}

		status := ""
		if len(list) < quota {
			status = "Accepted"
			fmt.Println("Status set automatically: Accepted")
		} else {
			fmt.Print("Quota full! Enter status (1 for Waiting List, 2 for Rejected): ")
			var choice int
			fmt.Scan(&choice)
			if choice == 1 {
				status = "Waiting List"
			} else {
				status = "Rejected"
			}
		}

		list = append(list, Participant{name, status})
	}

	fmt.Println("\n--- Final Summary ---")
	for _, p := range list {
		fmt.Printf("%s - %s\n", p.Name, p.Status)
	}
}
