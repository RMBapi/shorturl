package models

import (
	"strings"

	"example.com/shorturl/db"
	"example.com/shorturl/utils"
)

type URL struct {
	Id        int64
	LongURL   string `binding:"required"`
	UniqueKey string `binding:"required"`
	Shorturl  string `binding:"required"`
}

func (u *URL) UrlManager() error {

	key, err := utils.ProvideUniqueKey()

	if err != nil {
		return err
	}

	genarateKey := "https://ShortUrl/" + key
	u.UniqueKey = key
	u.Shorturl = genarateKey

	query := `INSERT INTO urllist(url,uniqueKey,shorturl)
	          VALUES(?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.LongURL, u.UniqueKey, u.Shorturl)
	if err != nil {
		return err
	}
	task_id, err := result.LastInsertId()
	u.Id = task_id
	return err
}

func Urlretrive(shortkey string) (string, error) {
	query := "SELECT url FROM urllist WHERE uniqueKey = ?"
	prefix := "https://ShortUrl/"
	uniqueKey := strings.TrimPrefix(shortkey, prefix)
	row := db.DB.QueryRow(query, uniqueKey)
	var l_key string
	err := row.Scan(&l_key)
	if err != nil {
		return "", err
	}
	return l_key, nil
}
