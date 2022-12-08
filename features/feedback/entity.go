package feedback

type CoreFeedback struct {
	ID             uint
	MenteeID       uint
	UserID         uint
	Description    string
	Mentee_Status  string
	Changed_Status string
	Mentees        CoreMentee
	Users          CoreUser
}

type CoreUser struct {
	ID        uint
	Full_Name string
}

type CoreMentee struct {
	ID          uint
	Mentee_Name string
}

type ServiceInterface interface {
	GetAll() (data []CoreFeedback, err error)
	Create(input CoreFeedback) (err error)
	GetById(id int) (data CoreFeedback, err error)
	UpdateFeedback(input CoreFeedback, id int) (err error)
	DeleteFeedback(id int) (err error)
}

type RepositoryInterface interface {
	GetAll() (data []CoreFeedback, err error)
	Create(input CoreFeedback) (row int, err error)
	GetById(id int) (data CoreFeedback, err error)
	UpdateFeedback(input CoreFeedback, id int) (err error)
	DeleteFeedback(id int) (row int, err error)
}
