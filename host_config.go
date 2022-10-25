package gopkhostmgmt

import (
	"os"

	"github.com/go-yaml/yaml"
)

type PkHostConfig struct {
	Name string `yaml:"name"`
}

type PkHostUser struct {
	Name         string `yaml:"name"`
	PasswordHash string `yaml:"password"`
	SshKey       string `yaml:"key"`
}

type PlaykeyConfig struct {
	Hostname       string   `yaml:"name"`
	VmTemplate     string   `yaml:"template"`
	SystemSnapshot string   `yaml:"system_snapshot"`
	HostCpus       int      `yaml:"host_cpus"`
	HostMemory     int      `yaml:"host_mem"`
	Vms            []HostVm `yaml:"vms"`
}

type HostVm struct {
	Name      string `yaml:"name"`
	IpAddress string `yaml:"ip_addr"`
	Cpu       int    `yaml:"num_cpus"`
	Memory    int    `yaml:"ram"`
	GpuAddr   string `yaml:"gpu_addr"`
}

func NewConfigFromFile(configPath string) (*PkHostConfig, error) {
	config := &PkHostConfig{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
