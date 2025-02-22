package models

import (
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

	genarateKey := "ShortUrl/" + key
	u.UniqueKey = key
	u.Shorturl = genarateKey

	query := `INSERT INTO urllist(url,uniqueKey,shorturl)
	          VALUES(?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// fmt.Println(u.LongURL)
	// fmt.Println(u.UniqueKey)
	// fmt.Println(u.Shorturl)

	result, err := stmt.Exec(u.LongURL, u.UniqueKey, u.Shorturl)
	if err != nil {
		return err
	}
	task_id, err := result.LastInsertId()
	u.Id = task_id
	return err
}
