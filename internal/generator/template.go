package generator

const (
	ZABBIX_ITEM_TYPE_DEPENDENT  ZabbixItemType = "DEPENDENT"
	ZABBIX_ITEM_TYPE_HTTP_AGENT ZabbixItemType = "HTTP_AGENT"

	ZABBIX_WIDGET_TYPE_GRAPH ZabbixWidgetType = "graph"

	ZABBIX_FIELD_TYPE_GRAPH           ZabbixFieldType = "GRAPH"
	ZABBIX_FIELD_TYPE_GRAPH_PROTOTYPE ZabbixFieldType = "GRAPH_PROTOTYPE"
	ZABBIX_FIELD_TYPE_STRING          ZabbixFieldType = "STRING"
	ZABBIX_FIELD_TYPE_INTEGER         ZabbixFieldType = "INTEGER"
	ZABBIX_FIELD_TYPE_HOST_GROUP      ZabbixFieldType = "HOST_GROUP"
	ZABBIX_FIELD_TYPE_HOST            ZabbixFieldType = "HOST"
	ZABBIX_FIELD_TYPE_ITEM            ZabbixFieldType = "ITEM"
	ZABBIX_FIELD_TYPE_ITEM_PROTOTYPE  ZabbixFieldType = "ITEM_PROTOTYPE"
	ZABBIX_FIELD_TYPE_MAP             ZabbixFieldType = "MAP"
	ZABBIX_FIELD_TYPE_SERVICE         ZabbixFieldType = "SERVICE"
	ZABBIX_FIELD_TYPE_SLA             ZabbixFieldType = "SLA"
	ZABBIX_FIELD_TYPE_USER            ZabbixFieldType = "USER"
	ZABBIX_FIELD_TYPE_ACTION          ZabbixFieldType = "ACTION"
	ZABBIX_FIELD_TYPE_MEDIA_TYPE      ZabbixFieldType = "MEDIA_TYPE"
)

var (
	// TODO: Create ZABBIX_WIDGET_TYPES from this list, as yet unused
	widgetTypes = []string{
		"actionlog",      // Action log
		"clock",          // Clock
		"discovery",      // Discovery status
		"favgraphs",      // Favorite graphs
		"favmaps",        // Favorite maps
		"gauge",          // Gauge
		"geomap",         // Geomap
		"graph",          // Graph (classic)
		"graphprototype", // Graph prototype
		"honeycomb",      // Honeycomb
		"hostavail",      // Host availability
		"hostcard",       // Host card
		"hostnavigator",  // Host navigator
		"itemhistory",    // Item history
		"itemnavigator",  // Item navigator
		"item",           // Item value
		"map",            // Map
		"navtree",        // Map Navigation Tree
		"piechart",       // Pie chart
		"problemhosts",   // Problem hosts
		"problems",       // Problems
		"problemsbysv",   // Problems by severity
		"slareport",      // SLA report
		"svggraph",       // Graph
		"systeminfo",     // System information
		"tophosts",       // Top hosts
		"topitems",       // Top items
		"toptriggers",    // Top triggers
		"trigover",       // Trigger overview
		"url",            // URL
		"web",            // Web monitoring
	}

	objectTypeLabels = map[ZabbixFieldType]string{}

	hostFieldTypes = []ZabbixFieldType{
		ZABBIX_FIELD_TYPE_GRAPH,
		ZABBIX_FIELD_TYPE_GRAPH_PROTOTYPE,
	}
	simpleFieldTypes = []ZabbixFieldType{
		ZABBIX_FIELD_TYPE_STRING,
		ZABBIX_FIELD_TYPE_INTEGER,
	}
)

type ZabbixFieldType string

type ZabbixItemType string

type ZabbixWidgetType string

func (ft ZabbixFieldType) String() string {
	label, ok := objectTypeLabels[ft]
	if !ok {
		return "Unknown"
	}

	return label
}

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

type Vendor struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Group struct {
	Name string `json:"name"`
}

type Graph struct {
	UUID       string      `json:"uuid"`
	Name       string      `json:"name"`
	GraphItems []GraphItem `json:"graph_items"`
}

type Item interface {
	ValidateType() bool
}

type DependentItem struct {
	UUID          string              `json:"uuid"`
	Name          string              `json:"name"`
	Type          ZabbixItemType      `json:"type"`
	Key           string              `json:"key"`
	Delay         string              `json:"delay"`
	ValueType     string              `json:"value_type"`
	Units         string              `json:"units"`
	Description   string              `json:"description"`
	PreProcessing []PreprocessingStep `json:"preprocessing"`
	MasterItem    MasterItem          `json:"master_item"`
	Tags          []Tag               `json:"tags"`
}

func (i DependentItem) ValidateType() bool {
	return i.Type == ZABBIX_ITEM_TYPE_DEPENDENT
}

type HTTPAgentItem struct {
	UUID          string              `json:"uuid"`
	Name          string              `json:"name"`
	Type          ZabbixItemType      `json:"type"`
	Key           string              `json:"key"`
	History       string              `json:"history"`
	Trends        string              `json:"trends"`
	ValueType     string              `json:"value_type"`
	Description   string              `json:"description"`
	PreProcessing []PreprocessingStep `json:"preprocessing"`
	URL           string              `json:"url"`
	RetrieveMode  string              `json:"retrieve_mode"`
	Tags          []Tag               `json:"tags"`
	Triggers      []Trigger           `json:"triggers"`
}

func (i HTTPAgentItem) ValidateType() bool {
	return i.Type == ZABBIX_ITEM_TYPE_HTTP_AGENT
}

// TODO: Flesh out this type
type DiscoveryRule struct {
}

type Macro struct {
	Macro       string `json:"macro"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
}

type Tag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Dashboard struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Pages []Page `json:"pages"`
}

type ValueMap struct {
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Mappings []Mapping `json:"mappings"`
}
type Page struct {
	Name    string   `json:"name"`
	Widgets []Widget `json:"widgets"`
}

type Widget struct {
	Name   string           `json:"name"`
	Type   ZabbixWidgetType `json:"type"`
	Width  string           `json:"width"`
	Height string           `json:"height"`
	X      string           `json:"x"`
	Y      string           `json:"y"`
	Fields []WidgetField    `json:"fields"`
}

type WidgetField interface {
	ValidateType() bool
}

type HostField struct {
	Type  ZabbixFieldType `json:"type"`
	Name  string          `json:"name"`
	Value HostFieldValue  `json:"value"`
}

func (f HostField) ValidateType() bool {
	for _, field := range hostFieldTypes {
		if field == f.Type {
			return true
		}
	}

	return false
}

type HostFieldValue struct {
	Host string `json:"host"`
	Name string `json:"name"`
}

type SimpleField struct {
	Type  ZabbixFieldType `json:"type"`
	Name  string          `json:"name"`
	Value string          `json:"value"`
}

func (f SimpleField) ValidateType() bool {
	for _, field := range simpleFieldTypes {
		if field == f.Type {
			return true
		}
	}

	return false
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

type Mapping struct {
	Value    string `json:"value"`
	NewValue string `json:"newvalue"`
}
type MasterItem struct {
	Key string `json:"key"`
}

type PreprocessingStep struct {
	Type         string   `json:"type"`
	Parameters   []string `json:"parameters"`
	ErrorHandler string   `json:"error_handler,omitempty"`
}

type Trigger struct {
	UUID        string `json:"uuid"`
	Expression  string `json:"expression"`
	Name        string `json:"name"`
	EventName   string `json:"event_name"`
	Priority    string `json:"priority"` // TODO: Make Type
	Description string `json:"description,omitempty"`
	ManualClose string `json:"manual_close,omitempty"`
	// Dependencies []Dependency `json:"dependencies"` //TODO: Make Type
	Tags []Tag `json:"tags,omitempty"`
}

type ValueMapRef struct {
	Name string `json:"name"`
}
