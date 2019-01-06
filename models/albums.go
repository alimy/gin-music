package models

type Album struct {
	Id          string
	Title       string
	Artist      string
	ReleaseYear string
	Genre       string
	TrackCount  int
	AlbumId     string
}

type Albums struct {
	AlbumSlice []Album
}
