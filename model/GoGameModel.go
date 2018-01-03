package model

//GameInfo struct type for a game information
type GameInfo struct {
	Name string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
}

//GameRelease data from the release api
type GameRelease struct {
	ID int `json:"id"`
	DataString string `json:"human"`
}

//GameData data from the release api
type GameData struct {
	Name string `json:"name"`
}