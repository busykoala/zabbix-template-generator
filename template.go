package zabbix_template_generator

type Template struct {
	UUID           string          `json:"uuid"`
	Host           string          `json:"host"`
	Template       string          `json:"template"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Dashboards     []Dashboard     `json:"dashboards"`
	DiscoveryRules []DiscoveryRule `json:"discovery_rules"`
	Groups         []Group         `json:"groups"`
	Items          []Item          `json:"items"`
	Macros         []Macro         `json:"macros"`
	Tags           []Tag           `json:"tags"`
	ValueMaps      []ValueMap      `json:"value_maps"`
}
