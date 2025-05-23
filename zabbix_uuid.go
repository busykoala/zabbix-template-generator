package zabbix_template_generator

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
