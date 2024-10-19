package proxmox

type Ticket struct {
	Data struct {
		Ticket              string `json:"ticket"`
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
		Username            string `json:"username"`
		Clustername         string `json:"clustername,omitempty"`
	} `json:"data"`
}
