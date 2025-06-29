package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

func PrintActiveProcesses() {}

func main() {
	fmt.Println("Current acive processes:")
	fmt.Printf("\n%-10s %-40s %-15s %-15s\n\n", "PId", "Process name", "CPU %", "Memory %")
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Can't get processes list:", err)
		return
	}
	for _, process := range processes {
		pid := process.Pid
		name, err := process.Name()
		if err != nil {
			continue
		}
		cpu, _ := process.CPUPercent()
		mem, _ := process.MemoryPercent()
		fmt.Printf("%-10d %-40s %-15.3f %-15.3f\n", pid, name, cpu, mem)
	}
}