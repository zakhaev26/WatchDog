package processor

import (
	"log"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

var cpuDetails []cpu.InfoStat

func FetchCpuUsage() (float64, string) {

	duration := 2 * time.Second
	interval := time.Second
	numSamples := int(duration / interval)

	var totalCPUUsage float64

	for i := 0; i < numSamples; i++ {
		cpuPercentages, err := cpu.Percent(interval, false)
		if err != nil {
			log.Fatal(err)
		}

		cpuUsage := cpuPercentages[0]
		totalCPUUsage += cpuUsage
		time.Sleep(interval)
	}

	// Calculate average CPU usage
	averageCPUUsage := totalCPUUsage / float64(numSamples)
	return averageCPUUsage, time.Now().Format("15:04:05")
}

func CpuUsageEachSecond() string {

	interval := time.Second

	cpuPercentages, err := cpu.Percent(interval, false)
	if err != nil {
		log.Fatal(err)
	}

	cpuUsage := cpuPercentages[0]
	data := strconv.FormatFloat(float64(cpuUsage), 'f', -1, 64)
	var res string = "CPU USAGE:" + data + " | TIME:" + time.Now().String()
	return res
}
