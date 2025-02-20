package hardware

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

const megaByteDiv uint64 = 1024 * 1024

func GetSystemSection() (string, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	hostStat, err := host.Info()
	if err != nil {
		return "", err
	}

	html := "<div class='system-data'><table class='table table-striped table-hover table-sm'><tbody>"
	html = html + "<tr><td>Platform:</td><td> " + hostStat.Platform + "</td></tr>"
	html = html + "<tr><td>Host name:</td><td>" + hostStat.Hostname + "</td></tr>"
	html = html + "<tr><td>Total memory:</td><td>" + strconv.FormatUint(vmStat.Total/megaByteDiv, 10) + " MB</td></tr>"
	html = html + "<tr><td>Free memory:</td><td>" + strconv.FormatUint(vmStat.Free/megaByteDiv, 10) + " MB</td></tr></tbody></table>"
	html = html + "</div>"

	return html, nil

}

func GetCpuSection() (string, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting CPU info", err)

	}
	percentage, err := cpu.Percent(0, true)
	if err != nil {
		return "", err
	}

	html := "<div class='cpu-data'><table class='table table-striped table-hover table-sm'><tbody>"

	if len(cpuStat) != 0 {
		html = html + "<tr><td>Model Name:</td><td>" + cpuStat[0].ModelName + "</td></tr>"
		html = html + "<tr style='display:none;'><td id='chart-data-dynamic'>"
		for idx, cpupercent := range percentage {
			if idx > 0 {
				html += ","
			}
			html = html + strconv.FormatFloat(cpupercent, 'f', 2, 64)
		}
		html = html + "</td></tr>"
	}

	firstCpus := percentage[:len(percentage)/2]
	secondCpus := percentage[len(percentage)/2:]

	html = html + "<tr><td>Cores: </td><td><div class='row mb-4'><div class='col-md-6'><table class='table table-sm'><tbody>"
	for idx, cpupercent := range firstCpus {
		html = html + "<tr><td>CPU [" + strconv.Itoa(idx) + "]: " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%</td></tr>"
	}
	html = html + "</tbody></table></div><div class='col-md-6'><table class='table table-sm'><tbody>"
	for idx, cpupercent := range secondCpus {
		html = html + "<tr><td>CPU [" + strconv.Itoa(idx+len(firstCpus)) + "]: " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%</td></tr>"
	}
	html = html + "</tbody></table></div></div></td></tr></tbody></table></div>"

	return html, nil
}

func GetCpuBalanceSection() (string, error) {
	_, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting CPU info", err)

	}
	percentage, err := cpu.Percent(0, true)
	if err != nil {
		return "", err
	}

	println(percentage)

	labels := ""
	data := ""
	for idx, cpupercent := range percentage {
		labels += fmt.Sprintf("'CPU %d', ", idx)
		data += fmt.Sprintf("%f, ", cpupercent)
	}

	html := ""

	return html, nil
}
