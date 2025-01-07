package store

import "strings"

func sortOrder(sortOrder string) string {
	sortOrder = strings.ToUpper(sortOrder)

	if sortOrder != "ASC" && sortOrder != "DESC" {
		return "ASC"
	}

	return sortOrder
}