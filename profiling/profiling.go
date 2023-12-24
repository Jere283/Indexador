package prof

import (
	"os"
	"runtime/pprof"
)

func StartCPUProfile() *os.File {
	cpuProf, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(cpuProf)
	return cpuProf
}

func StopCpuProfile(cpuProfile *os.File) {
	pprof.StopCPUProfile()
	cpuProfile.Close()
}

func StartMemoryProfile() *os.File {
	memoryProfile, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}
	return memoryProfile
}

func StopMemoryProfile(memoryProfile *os.File) {
	pprof.WriteHeapProfile(memoryProfile)
	memoryProfile.Close()
}
