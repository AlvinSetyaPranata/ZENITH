package master

type FacultyRequestModel struct {
	Name string `json:"name" validate:"required"`
}

type FacultyResponseModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DateCreated string `json:"created_at"`
}
