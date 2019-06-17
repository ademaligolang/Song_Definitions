package song

import "fmt"

type Name struct {
	Name string `json:"name"`
}

type Song struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Artists []Name `json:"artists"`
}

func (n *Song) String() string {
	return fmt.Sprintf("\tID: %d\n\tSong Title: %s\n\tArtists: %s\n", n.Id, n.Title, n.Artists)
}

type Songs struct {
	Songs []Song `json:"songs"`
}
