package interfaces

type AnimeStruct struct {
	AnimeId  int    `json:"anime_id"`
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
	Studio   string `json:"studio"`
}
