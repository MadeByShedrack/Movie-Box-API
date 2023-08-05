package models

type Movies struct {
	MovieID          uint   `json:"id" gorm:"primary_key"`
	MovieTitle       string `json:"title"`
	MovieGenre       string `json:"genre"`
	ReleaseDate      uint   `json:"releasedate"`
	Length           uint   `json:"length"`
	BoxOfficeRevenue uint   `json:"revenue"`
}
