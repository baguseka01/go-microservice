package paginator

type Meta struct {
	Page         int  `json:"page"`
	RowPerPage   int  `json:"row_per_page"`
	NextPage     bool `json:"next_page"`
	PreviousPage bool `json:"previous_page"`
}

func (meta *Meta) BuildMeta(dataLenght int, page int, rowPerPage int) {

	rowPerPage = rowPerPage - 1

	meta.Page = page
	meta.RowPerPage = rowPerPage
	meta.NextPage = false

	if dataLenght > rowPerPage {
		meta.NextPage = true
	}

	if (dataLenght-1 <= rowPerPage) && (page != 1) {
		meta.PreviousPage = true
	}
}