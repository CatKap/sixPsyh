package models

import "time"

//type User struct {
//    ID        int64     `json:"id"`
//    /Email     string    `json:"email"`
//    /Name      string    `json:"name"`
//    CreatedAt time.Time `json:"created_at"`
//}

type Meeting struct {
    ID        int64     `json:"id"`
		Cathegory string    `json:"cathegory"`
    Name      string    `json:"name"`
    description      string    `json:"desc"`
    CreatedAt time.Time `json:"created_at"`
    SceduledAt time.Time `json:"sceduled_at"`
}
