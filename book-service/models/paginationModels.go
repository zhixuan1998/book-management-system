package models

type Pagination struct {
	CurrentPage int  `json:"currentPage"`
	ItemFrom    int  `json:"itemFrom"`
	ItemTo      int  `json:"itemTo"`
	TotalItems  int  `json:"totalItems"`
	IsLastPage  bool `json:"isLastPage"`
}
