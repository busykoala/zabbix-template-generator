package generator

const (
	ZABBIX_WIDGET_FIELD_TYPE_INTEGER         ZabbixWidgetFieldType = "INTEGER"
	ZABBIX_WIDGET_FIELD_TYPE_STRING          ZabbixWidgetFieldType = "STRING"
	ZABBIX_WIDGET_FIELD_TYPE_HOST_GROUP      ZabbixWidgetFieldType = "HOST_GROUP"
	ZABBIX_WIDGET_FIELD_TYPE_HOST            ZabbixWidgetFieldType = "HOST"
	ZABBIX_WIDGET_FIELD_TYPE_ITEM            ZabbixWidgetFieldType = "ITEM"
	ZABBIX_WIDGET_FIELD_TYPE_ITEM_PROTOTYPE  ZabbixWidgetFieldType = "ITEM_PROTOTYPE"
	ZABBIX_WIDGET_FIELD_TYPE_GRAPH           ZabbixWidgetFieldType = "GRAPH"
	ZABBIX_WIDGET_FIELD_TYPE_GRAPH_PROTOTYPE ZabbixWidgetFieldType = "GRAPH_PROTOTYPE"
	ZABBIX_WIDGET_FIELD_TYPE_MAP             ZabbixWidgetFieldType = "MAP"
	ZABBIX_WIDGET_FIELD_TYPE_SERVICE         ZabbixWidgetFieldType = "SERVICE"
	ZABBIX_WIDGET_FIELD_TYPE_SLA             ZabbixWidgetFieldType = "SLA"
	ZABBIX_WIDGET_FIELD_TYPE_USER            ZabbixWidgetFieldType = "USER"
	ZABBIX_WIDGET_FIELD_TYPE_ACTION          ZabbixWidgetFieldType = "ACTION"
	ZABBIX_WIDGET_FIELD_TYPE_MEDIA_TYPE      ZabbixWidgetFieldType = "MEDIA_TYPE"
)

type ZabbixWidgetFieldType string

type WidgetField struct {
	Type ZabbixWidgetFieldType `json:"type"`
	Name string                `json:"name"`
	// TODO: https://www.zabbix.com/documentation/current/en/manual/api/reference/dashboard/widget_fields
	// The values depend on the type
	Value string `json:"value"`
}
