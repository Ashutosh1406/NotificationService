package models

type User struct {
	ID               string
	Name             string
	SubscribedGenres []string
}

func (u *User) SubsribeGenre(genreID string) {
	u.SubscribedGenres = append(u.SubscribedGenres, genreID)
}
