package paginator

import "strconv"

// CreatePagination return (skip, page, rowPerPage) will be build pagination variable , used for paginated
func CreatePagination(page string, rowPerPage string) (int, int, int) {

	pageValue, err := strconv.Atoi(page)
	if err != nil || pageValue <= 0 {
		pageValue = 1
	}

	rowPerPageValue, err := strconv.Atoi(rowPerPage)
	if err != nil || rowPerPageValue <= 0 {
		rowPerPageValue = 10
	}

	skip := (pageValue * rowPerPageValue) - rowPerPageValue

	rowPerPageValue = rowPerPageValue + 1

	return skip, pageValue, rowPerPageValue
}