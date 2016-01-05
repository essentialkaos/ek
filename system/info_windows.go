// +build !linux, !darwin, windows

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// LoadAvg contains information about average system load
type LoadAvg struct {
	Min1  float64 `json:"min1"`  // LA in last 1 minute
	Min5  float64 `json:"min5"`  // LA in last 5 minutes
	Min15 float64 `json:"min15"` // LA in last 15 minutes
	RProc int     `json:"rproc"` // Number of currently runnable kernel scheduling entities
	TProc int     `json:"tproc"` // Number of kernel scheduling entities that currently exist on the system
}

// MemInfo contains info about system memory
type MemInfo struct {
	MemTotal   uint64 `json:"total"`       // Total usable ram (i.e. physical ram minus a few reserved bits and the kernel binary code)
	MemFree    uint64 `json:"free"`        // The sum of MemFree - (Buffers + Cached)
	MemUsed    uint64 `json:"used"`        // MemTotal - MemFree
	Buffers    uint64 `json:"buffers"`     // Relatively temporary storage for raw disk blocks shouldn't get tremendously large (20MB or so)
	Cached     uint64 `json:"cached"`      // In-memory cache for files read from the disk (the pagecache).  Doesn't include SwapCached
	Active     uint64 `json:"active"`      // Memory that has been used more recently and usually not reclaimed unless absolutely necessary
	Inactive   uint64 `json:"inactive"`    // Memory which has been less recently used.  It is more eligible to be reclaimed for other purposes
	SwapTotal  uint64 `json:"swap_total"`  // Total amount of swap space available
	SwapFree   uint64 `json:"swap_free"`   // Memory which has been evicted from RAM, and is temporarily on the disk still also is in the swapfile
	SwapUsed   uint64 `json:"swap_used"`   // SwapTotal - SwapFree
	SwapCached uint64 `json:"spaw_cached"` // Memory that once was swapped out, is swapped back in but
	Dirty      uint64 `json:"dirty"`       // Memory which is waiting to get written back to the disk
	Slab       uint64 `json:"slab"`        // In-kernel data structures cache
}

// CPUInfo contains info about CPU usage
type CPUInfo struct {
	User   float64 `json:"user"`   // Normal processes executing in user mode
	System float64 `json:"system"` // Processes executing in kernel mode
	Nice   float64 `json:"nice"`   // Niced processes executing in user mode
	Idle   float64 `json:"idle"`   // Twiddling thumbs
	Wait   float64 `json:"wait"`   // Waiting for I/O to complete
	Count  int     `json:"count"`  // Number of CPU cores
}

// FSInfo contains info about fs usage
type FSInfo struct {
	Type    string   `json:"type"`    // FS type (ext4/ntfs/etc...)
	Device  string   `json:"device"`  // Device spec
	Used    uint64   `json:"used"`    // Used space
	Free    uint64   `json:"free"`    // Free space
	Total   uint64   `json:"total"`   // Total space
	IOStats *IOStats `json:"iostats"` // IO statistics
}

// IOStats contains inforamtion about I/O
type IOStats struct {
	ReadComplete  uint64 `json:"read_complete"`  // Reads completed successfully
	ReadMerged    uint64 `json:"read_merged"`    // Reads merged
	ReadSectors   uint64 `json:"read_sectors"`   // Sectors read
	ReadMs        uint64 `json:"read_ms"`        // Time spent reading (ms)
	WriteComplete uint64 `json:"write_complete"` // Writes completed
	WriteMerged   uint64 `json:"write_merged"`   // Writes merged
	WriteSectors  uint64 `json:"write_sectors"`  // Sectors written
	WriteMs       uint64 `json:"write_ms"`       // Time spent writing (ms)
	IOPending     uint64 `json:"io_pending"`     // I/Os currently in progress
	IOMs          uint64 `json:"io_ms"`          // Time spent doing I/Os (ms)
	IOQueueMs     uint64 `json:"io_queue_ms"`    // Weighted time spent doing I/Os (ms)
}

// SystemInfo contains info about system (hostname, OS, arch...)
type SystemInfo struct {
	Hostname string `json:"hostname"` // Hostname
	OS       string `json:"os"`       // OS name
	Kernel   string `json:"kernel"`   // Kernel version
	Arch     string `json:"arch"`     // System architecture (i386/i686/x86_64/etc...)
}

// InterfaceInfo contains info about network interfaces
type InterfaceInfo struct {
	ReceivedBytes      uint64 `json:"received_bytes"`
	ReceivedPackets    uint64 `json:"received_packets"`
	TransmittedBytes   uint64 `json:"transmitted_bytes"`
	TransmittedPackets uint64 `json:"transmitted_packets"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetUptime return uptime in seconds from 1/1/1970
func GetUptime() (uint64, error) {
	return 0, nil
}

// GetLA return loadavg
func GetLA() (*LoadAvg, error) {
	return &LoadAvg{}, nil
}

// GetMemInfo return memory info
func GetMemInfo() (*MemInfo, error) {
	return &MemInfo{}, nil
}

// GetCPUInfo return info about CPU usage
func GetCPUInfo() (*CPUInfo, error) {
	return &CPUInfo{}, nil
}

// GetFSInfo return info about mounted filesystems
func GetFSInfo() (map[string]*FSInfo, error) {
	return map[string]*FSInfo{"/": &FSInfo{}}, nil
}

// GetIOStats return I/O stats
func GetIOStats() (map[string]*IOStats, error) {
	return map[string]*IOStats{"/dev/sda1": &IOStats{}}, nil
}

// GetSystemInfo return system info
func GetSystemInfo() (*SystemInfo, error) {
	return &SystemInfo{}, nil
}

// GetInterfacesInfo return info about network interfaces
func GetInterfacesInfo() (map[string]*InterfaceInfo, error) {
	return map[string]*InterfaceInfo{"eth0": &InterfaceInfo{}}, nil
}

// GetNetworkSpeed return input/output speed in bytes per second
func GetNetworkSpeed() (uint64, uint64, error) {
	return 0, 0, nil
}

// GetIOUtil return IO utilization
func GetIOUtil() (map[string]float64, error) {
	return map[string]float64{"/": 0}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
