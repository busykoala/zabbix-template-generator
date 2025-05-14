package zabbix_template_generator

const (
	ZABBIX_ITEM_TYPE_DEPENDENT  ZabbixItemType = "DEPENDENT"
	ZABBIX_ITEM_TYPE_HTTP_AGENT ZabbixItemType = "HTTP_AGENT"
)

type Item interface {
	ValidateType() bool
}

type ZabbixItemType string

type DependentItem struct {
	UUID          string              `json:"uuid"`
	Name          string              `json:"name"`
	Type          ZabbixItemType      `json:"type"`
	Key           string              `json:"key"`
	Delay         string              `json:"delay,omitempty"`
	ValueType     ItemValueType       `json:"value_type"`
	Units         string              `json:"units,omitempty"`
	Description   string              `json:"description"`
	PreProcessing []PreprocessingStep `json:"preprocessing"`
	MasterItem    MasterItem          `json:"master_item"`
	Tags          []Tag               `json:"tags,omitempty"`
	Triggers      []Trigger           `json:"triggers,omitempty"`
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
	Trends        string              `json:"trends,omitempty"`
	ValueType     ItemValueType       `json:"value_type,omitempty"`
	Description   string              `json:"description"`
	PreProcessing []PreprocessingStep `json:"preprocessing,omitempty"`
	URL           string              `json:"url"`
	RetrieveMode  string              `json:"retrieve_mode,omitempty"`
	Tags          []Tag               `json:"tags,omitempty"`
	Triggers      []Trigger           `json:"triggers,omitempty"`
}

type ItemValueType string

const (
	ItemValueTypeUnsigned ItemValueType = "UNSIGNED"
	ItemValueTypeText     ItemValueType = "TEXT"
	ItemValueTypeFloat    ItemValueType = "FLOAT"
	ItemValueTypeChar     ItemValueType = "CHAR"
	ItemValueTypeLog      ItemValueType = "LOG"
	ItemValueTypeBinary   ItemValueType = "BINARY"
)

func (i HTTPAgentItem) ValidateType() bool {
	return i.Type == ZABBIX_ITEM_TYPE_HTTP_AGENT
}

type PreprocessingStep struct {
	Type         string   `json:"type"`
	Parameters   []string `json:"parameters"`
	ErrorHandler string   `json:"error_handler,omitempty"`
}

type MasterItem struct {
	Key string `json:"key"`
}

type Trigger struct {
	UUID        string          `json:"uuid"`
	Expression  string          `json:"expression"`
	Name        string          `json:"name"`
	EventName   string          `json:"event_name,omitempty"`
	Priority    TriggerPriority `json:"priority"`
	Description string          `json:"description,omitempty"`
	ManualClose string          `json:"manual_close,omitempty"`
	// Dependencies []Dependency `json:"dependencies"` // TODO: Make Type
	Tags []Tag `json:"tags,omitempty"`
}

type TriggerPriority string

const (
	PriorityInfo     TriggerPriority = "INFO"     // blue
	PriorityWarn     TriggerPriority = "WARNING"  // yellow
	PriorityAverage  TriggerPriority = "AVERAGE"  // orange
	PriorityHigh     TriggerPriority = "HIGH"     // red
	PriorityDisaster TriggerPriority = "DISASTER" // dark red
)
