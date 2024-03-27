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

	s := make([]string, n-1)
	slots := network.slots
	for i, v := range slots {
		if i != 0 {
			s[i-1] = strconv.Itoa(v)
		}
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
	devicesCnt[0] = k
	cntDevices := make(map[int][]int, n)
	cntDevices[k] = []int{0}
	for i := 1; i < n; i++ {
		cntDevices[0] = append(cntDevices[0], i)
	}
	cntDeviceList := make([]int, 0, n)
	cntDeviceList = append(cntDeviceList, 0)
	cntDeviceList = append(cntDeviceList, k)

	valuesTable := make([][]int, n)
	for i := 0; i < n; i++ {
		valuesTable[i] = make([]int, n)
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
		n:                n,
		k:                k,
	}
}

func (n *network) execute() {
	for n.totalUpdatesCnt != n.n {
		changes := make([][3]int, 0, 2*n.n)
		updates := n.findUpdates()
		for _, u := range updates {
			uploadDevices := n.findUploadDevices(u)
			device := n.findDownloadDevice(u)
			if len(uploadDevices) != 0 && device != -1 {
				change := n.handleDownloadDevice(device, uploadDevices, u)
				changes = append(changes, change)
				n.markedDownloaded[device] = true
			}
		}
		n.changeSlots()
		n.applyChanges(changes)
		n.recalculateCnts()
		n.reloadMarks()
		n.changeTotalUpdatesCnt()
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

func (n *network) findDeviceUpdate(device int) int {
	cntList := n.cntList
	for _, cnt := range cntList {
		cntUpdates := n.cntUpdates[cnt]
		for _, cntUpdate := range cntUpdates {
			if n.deviceUpdates[device][cntUpdate] {
				return cntUpdate
			}
		}
	}
	return -1
}

func (n *network) findUploadDevices(u int) []int {
	devices := make([]int, 0)
	cntDevicesList := n.cntDeviceList
	for _, cnt := range cntDevicesList {
		cntDevices := n.cntDevices[cnt]
		for _, v := range cntDevices {
			if !n.deviceUpdates[v][u] && !n.markedUploaded[v] {
				devices = append(devices, v)
			}
		}
	}
	return devices
}

func (n *network) findDownloadDevice(u int) int {
	cntDevicesList := n.cntDeviceList
	for _, cnt := range cntDevicesList {
		if cnt == 0 {
			continue
		}
		cntDevices := n.cntDevices[cnt]
		for _, v := range cntDevices {
			if n.deviceUpdates[v][u] && !n.markedDownloaded[v] {
				return v
			}
		}
	}
	return -1
}

func (n *network) handleDownloadDevice(device int, uploadDevices []int, update int) [3]int {
	value := -1
	uploadDevice := -1
	for _, v := range uploadDevices {
		if n.valuesTable[device][v] > value {
			value = n.valuesTable[device][v]
			uploadDevice = v
		}
		n.markedUploaded[v] = true
	}

	return [3]int{device, uploadDevice, update}
}

func (n *network) applyChanges(changes [][3]int) {
	for _, v := range changes {
		n.deviceUpdates[v[1]][v[2]] = true
		n.updatesCnt[v[2]]++
		n.devicesCnt[v[1]]++
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
	n.cntList = cntList

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
	n.cntDeviceList = cntDeviceList

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

func (n *network) changeTotalUpdatesCnt() {
	totalUpdatesCnt := 0
	for _, v := range n.devicesCnt {
		if v == n.k {
			totalUpdatesCnt++
		}
	}
	n.totalUpdatesCnt = totalUpdatesCnt
}

func (n *network) changeSlots() {
	for i, v := range n.devicesCnt {
		if v != n.k {
			n.slots[i]++
		}
	}
}
