package delivery

import "immersive-dashboard/features/mentee"

type MenteeResponse struct {
	ID                    uint		`json:"id"`
	Name                  string 	`json:"name"`
	ID_class              uint		`json:"id_class"`
	Status                string 	`json:"status"`
	Kategori              string 	`json:"kategori"`
	Gender                string 	`json:"gender"`
	Current_Address       string	`json:"current_address"`
	Home_Address          string 	`json:"home_address"`
	Email                 string 	`json:"email"`
	Telegram              string 	`json:"telegram"`
	Phone                 string 	`json:"phone"`
	Emergency_Name        string	`json:"emergency_name"`
	Emergency_Phone       string 	`json:"emergency-phone"`
	Emergency_Status      string 	`json:"emergency_status"`
	Education_Type        string 	`json:"education_type"`
	Education_Major       string 	`json:"education_major"`
	Education_Graduate    string 	`json:"education_graduate"`
	Education_Institution string	`json:"education_institution"`
}

func fromCore(dataCore user.Core) UserResponse {
	return UserResponse{
	ID                   datacore.ID,
	Name                 datacore.Name,
	ID_class             datacore.ID_Class,
	Status               datacore.Status,
	Kategori             datacore.Kategori,
	Gender               datacore.Gender,
	Current_Address      datacore.Current_Address,
	Home_Address         datacore.Home_Address,
	Email                datacore.Email,
	Telegram             datacore.Telegram,
	Phone                datacore.Phone,
	Emergency_Name       datacore.Emergency_Name,
	Emergency_Phone      datacore.Emergency_Phone,
	Emergency_Status     datacore.Emergency_Status,
	Education_Type       datacore.Education_Type,
	Education_Major      datacore.Education_Major,
	Education_Graduate   datacore.Education_Graduate,
	Education_Institution datacore.Education_Institution,
	}
}

func fromCoreList(dataCore []mentee.Core) []MenteeResponse {
	var dataResponse []MenteeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}