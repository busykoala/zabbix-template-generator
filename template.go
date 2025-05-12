package zabbix_template_generator

type Template struct {
	UUID           string          `json:"uuid"`
	Host           string          `json:"host,omitempty"`
	Template       string          `json:"template"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Dashboards     []Dashboard     `json:"dashboards,omitempty"`
	DiscoveryRules []DiscoveryRule `json:"discovery_rules,omitempty"`
	Groups         []Group         `json:"groups"`
	Items          []Item          `json:"items"`
	Macros         []Macro         `json:"macros,omitempty"`
	Tags           []Tag           `json:"tags,omitempty"`
	ValueMaps      []ValueMap      `json:"value_maps,omitempty"`
}
