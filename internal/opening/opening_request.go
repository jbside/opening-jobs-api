package opening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) TransformRequestToSchema() *Opening {
	return &Opening{
		Role:     r.Role,
		Company:  r.Company,
		Location: r.Location,
		Remote:   *r.Remote,
		Link:     r.Link,
		Salary:   r.Salary,
	}
}
