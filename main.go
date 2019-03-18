package main

import "github.com/loizoskounios/game-boy-emulator/cpu"

func main() {
	cpu := cpu.New()

	cpu.Dispatch()
}
