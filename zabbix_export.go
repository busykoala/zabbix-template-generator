package zabbix_template_generator

type ZabbixExport struct {
	Version        string          `json:"version"`
	TemplateGroups []TemplateGroup `json:"template_groups"`
	Templates      []Template      `json:"templates"`
	Graphs         []Graph         `json:"graphs"`
}

type TemplateGroup struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Vendor struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Graph struct {
	UUID       string      `json:"uuid"`
	Name       string      `json:"name"`
	GraphItems []GraphItem `json:"graph_items"`
}

type GraphItem struct {
	Color     string          `json:"color,omitempty"`
	SortOrder string          `json:"sortorder,omitempty"`
	Items     []GraphItemItem `json:"items"`
}

type GraphItemItem struct {
	Host string `json:"host"`
	Key  string `json:"key"`
}
