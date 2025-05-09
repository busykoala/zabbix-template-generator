package generator

type ZabbixWidgetType string

const (
	ZABBIX_WIDGET_TYPE_ACTIONLOG      ZabbixWidgetType = "actionlog"
	ZABBIX_WIDGET_TYPE_CLOCK          ZabbixWidgetType = "clock"
	ZABBIX_WIDGET_TYPE_DISCOVERY      ZabbixWidgetType = "discovery"
	ZABBIX_WIDGET_TYPE_FAVGRAPHS      ZabbixWidgetType = "favgraphs"
	ZABBIX_WIDGET_TYPE_FAVMAPS        ZabbixWidgetType = "favmaps"
	ZABBIX_WIDGET_TYPE_GAUGE          ZabbixWidgetType = "gauge"
	ZABBIX_WIDGET_TYPE_GEOMAP         ZabbixWidgetType = "geomap"
	ZABBIX_WIDGET_TYPE_GRAPH          ZabbixWidgetType = "graph"
	ZABBIX_WIDGET_TYPE_GRAPHPROTOTYPE ZabbixWidgetType = "graphprototype"
	ZABBIX_WIDGET_TYPE_HONEYCOMB      ZabbixWidgetType = "honeycomb"
	ZABBIX_WIDGET_TYPE_HOSTAVAIL      ZabbixWidgetType = "hostavail"
	ZABBIX_WIDGET_TYPE_HOSTCARD       ZabbixWidgetType = "hostcard"
	ZABBIX_WIDGET_TYPE_HOSTNAVIGATOR  ZabbixWidgetType = "hostnavigator"
	ZABBIX_WIDGET_TYPE_ITEMHISTORY    ZabbixWidgetType = "itemhistory"
	ZABBIX_WIDGET_TYPE_ITEMNAVIGATOR  ZabbixWidgetType = "itemnavigator"
	ZABBIX_WIDGET_TYPE_ITEM           ZabbixWidgetType = "item"
	ZABBIX_WIDGET_TYPE_MAP            ZabbixWidgetType = "map"
	ZABBIX_WIDGET_TYPE_NAVTREE        ZabbixWidgetType = "navtree"
	ZABBIX_WIDGET_TYPE_PIECHART       ZabbixWidgetType = "piechart"
	ZABBIX_WIDGET_TYPE_PROBLEMHOSTS   ZabbixWidgetType = "problemhosts"
	ZABBIX_WIDGET_TYPE_PROBLEMS       ZabbixWidgetType = "problems"
	ZABBIX_WIDGET_TYPE_PROBLEMSBYSV   ZabbixWidgetType = "problemsbysv"
	ZABBIX_WIDGET_TYPE_SLAREPORT      ZabbixWidgetType = "slareport"
	ZABBIX_WIDGET_TYPE_SVGGRAPH       ZabbixWidgetType = "svggraph"
	ZABBIX_WIDGET_TYPE_SYSTEMINFO     ZabbixWidgetType = "systeminfo"
	ZABBIX_WIDGET_TYPE_TOPHOSTS       ZabbixWidgetType = "tophosts"
	ZABBIX_WIDGET_TYPE_TOPITEMS       ZabbixWidgetType = "topitems"
	ZABBIX_WIDGET_TYPE_TOPTRIGGERS    ZabbixWidgetType = "toptriggers"
	ZABBIX_WIDGET_TYPE_TRIGOVER       ZabbixWidgetType = "trigover"
	ZABBIX_WIDGET_TYPE_URL            ZabbixWidgetType = "url"
	ZABBIX_WIDGET_TYPE_WEB            ZabbixWidgetType = "web"
)

type Dashboard struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Pages []Page `json:"pages"`
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
