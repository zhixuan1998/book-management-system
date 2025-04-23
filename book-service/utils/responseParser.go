package utils

import (
	"book-service/models"
	"slices"
)

func getSuccessStatus() models.Status {
	return models.Status{
		Code:    0,
		Message: "success",
	}
}

func getErrorStatus() models.Status {
	return models.Status{
		Code:    500,
		Message: "Internal server error.",
	}
}

func GenerateSuccessResponse() models.Response {
	return models.Response{
		Status: getSuccessStatus(),
	}
}

func GenerateSuccessResponseWithData(data any) models.ResponseWithData {
	return models.ResponseWithData{
		Status: getSuccessStatus(),
		Data:   data,
	}
}

func GenerateSuccessResponseWithPagination[T any](data []T, pagination models.Pagination) models.ResponseWithPagination {
	results := make([]any, len(data))

	for i, v := range slices.All(data) {
		results[i] = v
	}

	return models.ResponseWithPagination{
		Status: getSuccessStatus(),
		Data: models.DataWithPagination{
			Results:    results,
			Pagination: pagination,
		},
	}
}

func GenerateErrorResponse() models.ErrorResponse {
	return models.ErrorResponse{
		Status: getErrorStatus(),
	}
}
