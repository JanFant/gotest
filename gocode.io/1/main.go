// Objective: Set your gear to "full"

package main

func main() {
	agents := make([]Agent, 0)
	agents = append(agents, Agent{name: "J. Son", gear: "full"})
	agents = append(agents, Agent{name: "A. Pend", gear: "full"})
	agents = append(agents, Agent{name: "D. Buggs", gear: "none"})
	agents = append(agents, Agent{name: "X. Itwon", gear: "full"})
	agents = append(agents, Agent{name: "D. Fercloze", gear: "full"})

	// TODO: give "D. Buggs" full gear
	// NOTE: Output is validated, so don't use extra print statements
	// agents[2].gear = "full"
	for i := range agents {
		if agents[i].name == "D. Buggs" {
			agents[i].gear = "full"
			break
		}
	}

	println("Operation Go: Agent Manifest")
	println("----------------------------")
	for _, agent := range agents {
		println(agent.name, "-> Gear:", agent.gear)
	}
}

// Agent represents an agency employee
type Agent struct {
	name string
	gear string
}
