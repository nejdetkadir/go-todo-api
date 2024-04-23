package domain

import "math"

type GlobalErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type GlobalMessageResponse struct {
	Message string `json:"message"`
}

type PaginationMetaResponse struct {
	CurrentPage     int  `json:"current_page"`
	PrevPage        *int `json:"prev_page"`
	NextPage        *int `json:"next_page"`
	PerPage         int  `json:"per_page"`
	TotalPagesCount int  `json:"total_pages_count"`
	TotalCount      int  `json:"total_count"`
	IsFirstPage     bool `json:"is_first_page"`
	IsLastPage      bool `json:"is_last_page"`
	IsEmpty         bool `json:"is_empty"`
}

type PaginationRequest struct {
	Page    int `query:"page" validate:"numeric,min=1"`
	PerPage int `query:"per_page" validate:"numeric,min=1,max=100"`
}

func (p PaginationRequest) GetOffset() int {
	return (p.Page - 1) * p.PerPage
}

func (p PaginationRequest) GetLimit() int {
	return p.PerPage
}

func (pl PaginationMetaResponse) GetPaginationMetaResponse(paginationRequest PaginationRequest, totalItemsCount int, currentItemsCount int) PaginationMetaResponse {
	var currentPage = paginationRequest.Page
	var totalPagesCount int
	var prevPage *int
	var nextPage *int

	if totalItemsCount == 0 {
		totalPagesCount = 1
	} else {
		totalPagesCount = int(math.Ceil(float64(totalItemsCount) / float64(paginationRequest.PerPage)))
	}

	if currentPage-1 < 1 {
		prevPage = nil
	} else if currentPage-1 > totalPagesCount {
		prevPage = nil
	} else {
		var newPrevPage = currentPage - 1
		prevPage = &newPrevPage
	}

	if currentPage+1 > totalPagesCount {
		nextPage = nil
	} else {
		var newNextPage = currentPage + 1
		nextPage = &newNextPage
	}

	return PaginationMetaResponse{
		CurrentPage:     currentPage,
		PrevPage:        prevPage,
		NextPage:        nextPage,
		PerPage:         paginationRequest.PerPage,
		TotalPagesCount: totalPagesCount,
		TotalCount:      totalItemsCount,
		IsFirstPage:     paginationRequest.Page == 1,
		IsLastPage:      currentPage == totalPagesCount || totalPagesCount == 0,
		IsEmpty:         currentItemsCount == 0,
	}
}
