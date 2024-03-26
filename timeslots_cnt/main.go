package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type network struct {
	deviceUpdates    [][]bool
	n                int
	k                int
	deviceUpdatesCnt []int
	updatesCnt       []int
	cntUpdates       map[int][]int
	cntList          []int
	devicesCnt       []int
	cntDevices       map[int][]int
	cntDeviceList    []int
	valuesTable      [][]int
	markedUploaded   []bool
	markedDownloaded []bool
	totalUpdatesCnt  int
	slots            []int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	network := newNetwork(n, k)
	network.execute()

	s := make([]string, n)
	slots := network.slots
	for i, v := range slots {
		s[i] = strconv.Itoa(v)
	}

	fmt.Println(strings.Join(s, " "))
}

func newNetwork(n int, k int) *network {
	deviceUpdates := make([][]bool, n)
	for i := 0; i < n; i++ {
		deviceUpdates[i] = make([]bool, k)
	}
	for j := 0; j < k; j++ {
		deviceUpdates[0][j] = true
	}

	deviceUpdatesCnt := make([]int, n)
	deviceUpdatesCnt[0] = k

	updatesCnt := make([]int, k)
	cntUpdates := make(map[int][]int, k)
	cntList := make([]int, 0, k)
	for i := 0; i < k; i++ {
		updatesCnt[i] = 1
		if _, ok := cntUpdates[1]; !ok {
			cntList = append(cntList, 1)
		}
		cntUpdates[1] = append(cntUpdates[1], i)
	}

	devicesCnt := make([]int, n)
	devicesCnt[0] = n
	cntDevices := make(map[int][]int, n)
	cntDevices[n] = []int{0}
	cntDeviceList := make([]int, 0, n)
	cntDeviceList = append(cntDeviceList, n)

	valuesTable := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				valuesTable[i][j] = -1
			}
		}
	}

	markedDownloaded := make([]bool, n)
	markedUploaded := make([]bool, n)
	markedUploaded[0] = true

	slots := make([]int, n)

	return &network{
		deviceUpdates:    deviceUpdates,
		deviceUpdatesCnt: deviceUpdatesCnt,
		updatesCnt:       updatesCnt,
		cntUpdates:       cntUpdates,
		cntList:          cntList,
		devicesCnt:       devicesCnt,
		cntDevices:       cntDevices,
		cntDeviceList:    cntDeviceList,
		valuesTable:      valuesTable,
		markedDownloaded: markedDownloaded,
		markedUploaded:   markedUploaded,
		totalUpdatesCnt:  1,
		slots:            slots,
	}
}

func (n *network) execute() {
	for n.totalUpdatesCnt != n.n {
		updates := n.findUpdates()
		changes := make([][3]int, 0, 2*n.n)
		for _, update := range updates {
			downloadDevices := n.findDownloadDevices(update)
			changes = append(changes, n.handleDownloadDevices(downloadDevices, update)...)
		}
		n.applyChanges(changes)
		n.recalculateCnts()
		n.reloadMarks()
		n.changeTotalUpdatesCnt()
		n.changeSlots()
	}
}

func (n *network) findUpdates() []int {
	updates := make([]int, 0, n.k)
	cntList := n.cntList
	for _, cnt := range cntList {
		cntUpdates := n.cntUpdates[cnt]
		for _, cntUpdate := range cntUpdates {
			updates = append(updates, cntUpdate)
		}
	}
	return updates
}

func (n *network) findDownloadDevices(update int) []int {
	devices := make([]int, 0, n.n)
	deviceUpdates := n.deviceUpdates
	cntDevicesList := n.cntDeviceList
	for _, cnt := range cntDevicesList {
		cntDevices := n.cntDevices[cnt]
		for _, v := range cntDevices {
			if deviceUpdates[v][update] && !n.markedDownloaded[v] {
				devices = append(devices, v)
			}
		}
	}
	return devices
}

func (n *network) handleDownloadDevices(devices []int, update int) [][3]int {
	changes := make([][3]int, 0, n.n)
	for _, d := range devices {
		mostValuedDevice := n.findMostValuedDevice(d)
		if mostValuedDevice != -1 {
			n.markedDownloaded[d] = true
			n.markedUploaded[mostValuedDevice] = true
			changes = append(changes, [3]int{d, mostValuedDevice, update})
		}
	}
	return changes
}

func (n *network) findMostValuedDevice(device int) int {
	table := n.valuesTable
	list := make([]int, 0, n.n)
	startValueIndex := make(map[int]int, n.n)
	for i := 0; i < n.n; i++ {
		if table[device][i] != -1 && !n.markedUploaded[i] && n.deviceUpdatesCnt[i] != n.n {
			value := table[device][i]
			if _, ok := startValueIndex[value]; !ok {
				list = append(list, value)
				startValueIndex[value] = i
			}
		}
	}
	slices.Sort(list)

	listCnt := len(list)

	if listCnt != 0 {
		return startValueIndex[listCnt-1]
	}
	return -1
}

func (n *network) applyChanges(changes [][3]int) {
	for _, v := range changes {
		n.deviceUpdates[v[1]][v[2]] = true
		n.deviceUpdatesCnt[v[1]]++
		n.updatesCnt[v[2]]++
		n.devicesCnt[v[2]]++
		n.valuesTable[v[0]][v[1]]++
		n.valuesTable[v[1]][v[0]]++
	}
}

func (n *network) recalculateCnts() {
	for k := range n.cntUpdates {
		delete(n.cntUpdates, k)
	}
	n.cntList = n.cntList[:0]

	updatesCnt := n.updatesCnt
	cntUpdates := n.cntUpdates
	cntList := n.cntList
	for i := 0; i < n.k; i++ {
		cnt := updatesCnt[i]
		if _, ok := cntUpdates[cnt]; !ok {
			cntList = append(cntList, cnt)
		}
		cntUpdates[cnt] = append(cntUpdates[cnt], i)
	}

	slices.Sort(cntList)

	for k := range n.cntDevices {
		delete(n.cntDevices, k)
	}
	n.cntDeviceList = n.cntDeviceList[:0]

	devicesCnt := n.devicesCnt
	cntDevices := n.cntDevices
	cntDeviceList := n.cntDeviceList
	for i := 0; i < n.n; i++ {
		cnt := devicesCnt[i]
		if _, ok := cntDevices[cnt]; !ok {
			cntDeviceList = append(cntDeviceList, cnt)
		}
		cntDevices[cnt] = append(cntDevices[cnt], i)
	}

	slices.Sort(cntDeviceList)
}

func (n *network) reloadMarks() {
	for i := range n.markedUploaded {
		n.markedUploaded[i] = false
	}

	for i := range n.markedDownloaded {
		n.markedDownloaded[i] = false
	}
}

func (n *network) changeTotalUpdatesCnt() int {
	totalUpdatesCnt := 0
	for _, v := range n.deviceUpdatesCnt {
		if v == n.k {
			totalUpdatesCnt++
		}
	}
	return totalUpdatesCnt
}

func (n *network) changeSlots() {
	for i, v := range n.deviceUpdatesCnt {
		if v != n.k {
			n.slots[i]++
		}
	}
}
