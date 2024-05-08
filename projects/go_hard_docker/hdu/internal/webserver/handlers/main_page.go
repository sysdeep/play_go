package handlers

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types/system"
	"github.com/labstack/echo/v4"
)

// main page model
type mainPageModel struct {
	DaemonHost    string // DaemonHost returns the host address used by the client
	ClientVersion string // the API version used by this client
	SystemInfo    systemInfo
}

// main page handler
func (h *Handlers) MainPage(c echo.Context) error {

	// слишком долго
	// disk_usage, err := h.docker_client.DiskUsage(context.Background(), types.DiskUsageOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", disk_usage)

	sys_info, err := h.docker_client.Info(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	response := mainPageModel{
		DaemonHost:    h.docker_client.DaemonHost(),
		ClientVersion: h.docker_client.ClientVersion(),
		SystemInfo:    make_system_info(sys_info),
	}

	return c.Render(http.StatusOK, "main", response)
}

// systemInfo
type systemInfo struct {
	ID                string
	Containers        int
	ContainersRunning int
	ContainersPaused  int
	ContainersStopped int
	Images            int
	Driver            string
	// DriverStatus       [][2]string
	// SystemStatus       [][2]string `json:",omitempty"` // SystemStatus is only propagated by the Swarm standalone API
	// Plugins            PluginsInfo
	// MemoryLimit        bool
	// SwapLimit          bool
	// KernelMemory       bool `json:",omitempty"` // Deprecated: kernel 5.4 deprecated kmem.limit_in_bytes
	// KernelMemoryTCP    bool `json:",omitempty"` // KernelMemoryTCP is not supported on cgroups v2.
	// CPUCfsPeriod       bool `json:"CpuCfsPeriod"`
	// CPUCfsQuota        bool `json:"CpuCfsQuota"`
	// CPUShares          bool
	// CPUSet             bool
	// PidsLimit          bool
	// IPv4Forwarding     bool
	// BridgeNfIptables   bool
	// BridgeNfIP6tables  bool `json:"BridgeNfIp6tables"`
	// Debug              bool
	// NFd                int
	// OomKillDisable     bool
	// NGoroutines        int
	// SystemTime         string
	// LoggingDriver      string
	// CgroupDriver       string
	// CgroupVersion      string `json:",omitempty"`
	// NEventsListener    int
	// KernelVersion      string
	// OperatingSystem    string
	// OSVersion          string
	// OSType             string
	// Architecture       string
	// IndexServerAddress string
	// RegistryConfig     *registry.ServiceConfig
	// NCPU               int
	// MemTotal           int64
	// GenericResources   []swarm.GenericResource
	// DockerRootDir      string
	// HTTPProxy          string `json:"HttpProxy"`
	// HTTPSProxy         string `json:"HttpsProxy"`
	// NoProxy            string
	// Name               string
	// Labels             []string
	// ExperimentalBuild  bool
	// ServerVersion      string
	// Runtimes           map[string]RuntimeWithStatus
	// DefaultRuntime     string
	// Swarm              swarm.Info
	// // LiveRestoreEnabled determines whether containers should be kept
	// // running when the daemon is shutdown or upon daemon start if
	// // running containers are detected
	// LiveRestoreEnabled  bool
	// Isolation           container.Isolation
	// InitBinary          string
	// ContainerdCommit    Commit
	// RuncCommit          Commit
	// InitCommit          Commit
	// SecurityOptions     []string
	// ProductLicense      string               `json:",omitempty"`
	// DefaultAddressPools []NetworkAddressPool `json:",omitempty"`
	// CDISpecDirs         []string
	//
	// // Legacy API fields for older API versions.
	// legacyFields
	//
	// // Warnings contains a slice of warnings that occurred  while collecting
	// // system information. These warnings are intended to be informational
	// // messages for the user, and are not intended to be parsed / used for
	// // other purposes, as they do not have a fixed format.
	// Warnings []string
}

func make_system_info(data system.Info) systemInfo {
	return systemInfo{
		ID:                data.ID,
		Containers:        data.Containers,
		ContainersRunning: data.ContainersRunning,
		ContainersPaused:  data.ContainersPaused,
		ContainersStopped: data.ContainersStopped,
		Images:            data.Images,
		Driver:            data.Driver,
	}
}
