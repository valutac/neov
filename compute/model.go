package compute

type Flavor struct {
	Disk       int
	Ephemeral  int
	ExtraSpecs struct {
		CPUModel    string `json:"hw:cpu_model"`
		CPUPolicy   string `json:"hw:cpu_policy"`
		MEMPageSize string `json:"hw:mem_page_size"`
	}
	Name  string `json:"original_name"`
	RAM   int
	Swap  int
	VCPUs int
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type Server struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Links      []Link  `json:"links"`
	Flavor     *Flavor `json:"flavor,omitempty"`
	Status     string  `json:"status"`
	HostStatus string  `json:"host_status"`
}
