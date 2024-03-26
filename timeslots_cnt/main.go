package main

import "slices"

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
}

func main() {

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
	}
}

func (n *network) execute() {
	updates := n.findUpdates()
	changes := make([][3]int, 2*n.n)
	for _, update := range updates {
		downloadDevices := n.findDownloadDevices(update)
		changes = append(changes, n.handleDownloadDevices(downloadDevices, update)...)
	}
	n.applyChanges(changes)
}

func (n *network) findUpdates() []int {
	updates := make([]int, n.k)
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
			if deviceUpdates[v][update] {
				devices = append(devices, v)
			}
		}
	}
	return devices
}

func (n *network) handleDownloadDevices(devices []int, update int) [][3]int {
	changes := make([][3]int, 0, n.n)
	for _, d := range devices {
		changes = append(changes, [3]int{d, n.findMostValuedDevice(d), update})
	}
	return changes
}

func (n *network) findMostValuedDevice(device int) int {
	table := n.valuesTable
	list := make([]int, 0, n.n)
	startValueIndex := make(map[int]int, n.n)
	for i := 0; i < n.n; i++ {
		if i != -1 {
			value := table[device][i]
			if _, ok := startValueIndex[value]; !ok {
				list = append(list, value)
				startValueIndex[value] = i
			}
		}
	}
	slices.Sort(list)

	listCnt := len(list)
	return startValueIndex[listCnt-1]
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
