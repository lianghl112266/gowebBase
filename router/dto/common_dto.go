package dto

type CommIDDTO struct {
	ID uint `json:"id" from:"id" uri:"id"`
}

type Paginate struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (me *Paginate) GetPage() int {
	if me.Page <= 0 {
		me.Page = 1
	}
	return me.Page
}

func (me *Paginate) GetLimit() int {
	if me.Limit <= 0 {
		me.Limit = 10
	}
	return me.Limit
}
