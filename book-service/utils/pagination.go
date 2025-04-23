package utils

import "book-service/models"

func GeneratePagination(page, limit, count int) models.Pagination {

	offset := (page - 1) * limit
	from := offset + 1

	to := from - 1

	if limit > count {
		to += count
	} else {
		to += limit
	}

	return models.Pagination{
		CurrentPage: page,
		ItemFrom: from,
		ItemTo: to,
		TotalItems: count,
		IsLastPage: count <= limit,
	}
} 