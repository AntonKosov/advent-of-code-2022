package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2022/aoc"
)

const minutes = 26
const variantsThreshold = 1_000

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (rates []uint16, tunnels [][]tunnel, start byte) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	valves := make(map[string]byte, len(lines))
	for i, line := range lines {
		valve := line[6:8]
		valves[valve] = byte(i)
	}

	rates = make([]uint16, len(valves))
	tunnels = make([][]tunnel, len(valves))

	for valveID, line := range lines {
		parts := strings.Split(line[23:], " ")
		rate := aoc.StrToInt(parts[0][:len(parts[0])-1])
		var connectedTunnels []tunnel
		for i := 5; i < len(parts); i++ {
			c := parts[i]
			connectedTunnels = append(connectedTunnels, tunnel{to: valves[c[:2]], length: 1})
		}

		rates[valveID] = uint16(rate)
		tunnels[valveID] = connectedTunnels
	}

	return rates, tunnels, valves["AA"]
}

func process(rates []uint16, tunnels [][]tunnel, start byte) uint16 {
	reduceGraph(rates, tunnels, start)
	initValveMap(tunnels)
	var valves valveStatus
	if rates[start] == 0 {
		valves.open(start)
	}
	max := uint16(0)
	toProcess := []memo{{}}
	toProcess[0].add(state{
		players: [2]player{
			{tunnel: start, minutesLeft: minutes},
			{tunnel: start, minutesLeft: minutes},
		},
		valves: valves,
	})

	for processed := true; processed; {
		processed = false
		for i := 0; i < len(toProcess); i++ {
			states := toProcess[i]
			if len(states) == 0 {
				continue
			}
			processed = true
			toProcess[i] = memo{}
			nextStates := processStates(states, rates, tunnels)
			for s := range nextStates {
				i := s.players[0].minutesLeft + s.players[1].minutesLeft
				for int(i) >= len(toProcess) {
					toProcess = append(toProcess, memo{})
				}
				toProcess[i][s] = struct{}{}
				if s.pressure > max {
					max = s.pressure
					fmt.Println("max pressure:", max)
				}
			}
		}
	}

	return max
}

func processStates(states memo, rates []uint16, tunnels [][]tunnel) memo {
	discardPressure := calcDiscardLevel(states)
	nextToProcess := make(memo, len(states))
	for s := range states {
		if s.players[0].minutesLeft <= 1 && s.players[1].minutesLeft <= 1 {
			continue
		}

		if s.pressure < discardPressure {
			continue
		}

		player := 0
		if s.players[0].minutesLeft < s.players[1].minutesLeft {
			player = 1
		}

		if valve := s.players[player].tunnel; s.valves.isClosed(valve) {
			sc := s
			sc.valves.open(valve)
			sc.players[player].minutesLeft--
			sc.pressure += rates[valve] * uint16(sc.players[player].minutesLeft)
			nextToProcess.add(sc)
		}

		for _, nextTunnel := range tunnels[s.players[player].tunnel] {
			if s.players[player].minutesLeft < nextTunnel.length+1 {
				continue
			}
			sc := s
			sc.players[player].minutesLeft -= nextTunnel.length
			sc.players[player].tunnel = nextTunnel.to
			nextToProcess.add(sc)
		}
	}

	return nextToProcess
}

func calcDiscardLevel(m memo) uint16 {
	// the data must be grouped by the sum of left minutes
	if len(m) < variantsThreshold {
		return 0
	}
	totalPressure := 0
	for s := range m {
		totalPressure += int(s.pressure)
	}
	avr := totalPressure / len(m)
	return uint16(avr)
}

func reduceGraph(rates []uint16, tunnels [][]tunnel, start byte) {
	n := byte(len(rates))
	for tunnelToRemove := byte(0); tunnelToRemove < n; tunnelToRemove++ {
		if tunnelToRemove == start || rates[tunnelToRemove] != 0 {
			continue
		}

		for t := byte(0); t < n; t++ {
			if t == tunnelToRemove {
				continue
			}

			for i := byte(0); i < byte(len(tunnels[t])); {
				tunnelToShorten := tunnels[t][i]
				if tunnelToShorten.to != tunnelToRemove {
					i++
					continue
				}
				tunnels[t] = append(tunnels[t][:i], tunnels[t][i+1:]...)
				for _, rt := range tunnels[tunnelToRemove] {
					if rt.to == t {
						// back path
						continue
					}
					tunnels[t] = append(tunnels[t], tunnel{
						to:     rt.to,
						length: tunnelToShorten.length + rt.length,
					})
				}
			}
		}

		tunnels[tunnelToRemove] = nil
	}
}

type player struct {
	tunnel      byte
	minutesLeft byte
}

type state struct {
	players  [2]player
	pressure uint16
	valves   valveStatus
}

var valveMask [64]uint16 // graph valves to valveStatus
var activeValves []byte

func initValveMap(tunnels [][]tunnel) {
	i := uint16(1)
	count := 0
	for idx, t := range tunnels {
		if len(t) == 0 {
			continue
		}
		count++
		activeValves = append(activeValves, byte(idx))
		valveMask[idx] = i
		i <<= 1
	}
	if count > 16 {
		panic("the number of tunnels must be 16 max")
	}
}

type valveStatus uint16

func (vs *valveStatus) isClosed(valve byte) bool {
	vm := valveMask[valve]
	return uint16(*vs)&vm == 0
}

func (vs *valveStatus) open(valve byte) {
	(*vs) |= valveStatus(valveMask[valve])
}

type tunnel struct {
	to     byte
	length byte
}

type memo map[state]struct{}

func (m memo) add(s state) bool {
	if _, ok := m[s]; ok {
		return false
	}
	s.players[0], s.players[1] = s.players[1], s.players[0]
	if _, ok := m[s]; ok {
		return false
	}
	m[s] = struct{}{}
	return true
}
