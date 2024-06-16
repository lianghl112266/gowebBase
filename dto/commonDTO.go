package dto

// CommID represents a common ID structure used for various requests.
type CommID struct {
	ID uint `json:"id" from:"id" uri:"id"`
}

// Paginate represents pagination parameters for requests.
type Paginate struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

// GetPage returns the current page number. If the page is not set or less than 1, it defaults to 1.
func (me *Paginate) GetPage() int {
	if me.Page <= 0 {
		me.Page = 1
	}
	return me.Page
}

// GetLimit returns the limit of items per page. If the limit is not set or less than 1, it defaults to 10.
func (me *Paginate) GetLimit() int {
	if me.Limit <= 0 {
		me.Limit = 10
	}
	return me.Limit
}
