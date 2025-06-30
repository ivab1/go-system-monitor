package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

func PrintResourcesUsage() {
	cpus, err := cpu.Percent(time.Second, false)
	if err != nil {
		return
	}
	cpu := cpus[0]
	fmt.Printf("CPU usage: %.2f%%\n", cpu)

	memory, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	memUsage := float64(memory.Used) / 1024 / 1024 / 1024
	memTotal := float64(memory.Total) / 1024 / 1024 / 1024
	fmt.Printf("Memory usage: %.2f Gb of %.2f Gb\n", memUsage, memTotal)
}

func PrintActiveProcesses() {
	fmt.Println("\nCurrent acive processes:")
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

func main() {
	PrintResourcesUsage()
	PrintActiveProcesses()
}