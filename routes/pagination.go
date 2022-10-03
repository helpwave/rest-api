package routes

type PaginatedResponse struct {
	NumPages  uint
	Page      uint
	PageSize  uint
	TotalSite uint
	LastPage  bool
}
