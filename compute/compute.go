package compute

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
)

const ComputeUrl = "https://nova.wjv-1.neo.id:8774/v2"

type service struct {
	token string
}

func NewService(token string) *service {
	return &service{token}
}

func (srv *service) GetServers(projectID string) []Server {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " get server list..."
	s.Start()

	url := fmt.Sprintf("%s/%s/servers/detail", ComputeUrl, projectID)
	// fmt.Printf("get server list from: %s\n", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-Auth-Token", srv.token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var response struct {
		servers []Server
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Got %d Compute Engine\n", len(response.servers))
	s.Stop()
	return response.servers
}

func DisplayServerList(servers []Server) {
	data := [][]string{}
	for _, server := range servers {
		current := []string{server.ID, server.Name, server.Status, server.HostStatus}
		if server.Flavor != nil {
			current = append(current, strconv.Itoa(server.Flavor.RAM))
			current = append(current, strconv.Itoa(server.Flavor.VCPUs))
		}
		data = append(data, current)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Compute ID", "Name", "Status", "Host Status", "RAM", "vCPUs"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()

	fmt.Println()
	fmt.Printf("Total Server: %d\n", len(servers))
}
