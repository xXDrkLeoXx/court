package model

type Court struct {
	Id        uint
	LongLink  string
	ShortLink string
}

func CreateShortLink(link string) error {
	query := "INSERT INTO court(long_link, short_link) VALUES($1, $2)"
	randId := "/AAAAAA"
	_, err := db.Query(query, link, randId)
}

func CreateCustomLink(link, path string) {
	query := "INSERT INTO court(long_link, short_link) VALUES($1, $2)"
	_, err := db.Query(query, link, path)
}

func GetAllLinks() ([]Court, error) {
	links := []Court{}

	query := "SELECT id, long_link, short_link FROM court"
	rows, err := db.Query(query)
	if err != nil {
		return links, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		var shortLink, longLink string
		err := db.Scan(&id, &longLink, &shortLink)
		if err != nil {
			return links, err
		}
		link := Court{
			Id:        id,
			ShortLink: shortLink,
			LongLink:  longLink,
		}
		links = append(links, link)
	}
	return links, err
}

func GetLink(path string) (string, error) {
	var link string
	query := "SELECT long_link FROM court WHERE short_link=$1"
	row, err := db.Query(query, path)
	if err != nil {
		return link, err
	}
	for row.Next() {
		err := row.Scan(&link)
		if err != nil {
			return link, err
		}
	}
	return link, err
}

func Delete(id uint) error {
	query := "DELETE FROM court WHERE id=$1"
	_, err := db.Exec(query, id)
	return err
}
