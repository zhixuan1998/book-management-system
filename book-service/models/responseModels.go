package models

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataWithPagination struct {
	Results    []any `json:"results"`
	Pagination any   `json:"pagination"`
}

type Response struct {
	Status Status `json:"status"`
}

type ResponseWithData struct {
	Status Status `json:"status"`
	Data   any    `json:"data"`
}

type ResponseWithPagination struct {
	Status Status             `json:"status"`
	Data   DataWithPagination `json:"data"`
}

type ErrorResponse struct {
	Status Status `json:"status"`
}
