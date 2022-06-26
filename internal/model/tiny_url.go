package model

import (
	"errors"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(dsn string) error {
	tmp, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db = tmp
	return nil
}

type TinyUrl struct {
	Id        int    `json:"id" db:"id"`
	BaseUrl   string `db:"base_url"`
	SuffixUrl string `db:"suffix_url"`
	FullUrl   string `db:"full_url"`
	ShotCode  string `db:"shot_code"`
}

func CreateTinyUrl(oriUrlRaw, shot string) error {
	oriUrl, err := url.Parse(oriUrlRaw)
	if err != nil {
		return err
	}
	baseUrl := oriUrl.Scheme + "://" + oriUrl.Host
	suffixUrl := oriUrl.RequestURI()
	fmt.Println(baseUrl, suffixUrl)
	tx := db.MustBegin()
	_, err = tx.NamedExec("INSERT INTO tiny_url (base_url,suffix_url,full_url,shot_code) VALUES (:base_url,:suffix_url,:full_url,:shot_code)",
		&TinyUrl{BaseUrl: baseUrl, SuffixUrl: suffixUrl, FullUrl: oriUrlRaw, ShotCode: shot})
	if err != nil {
		return err
	}
	return tx.Commit()
}

func RetrieveByShotCode(shot string) (*TinyUrl, error) {
	if len(shot) != 6 {
		return nil, errors.New("invalid shot code")
	}
	result := TinyUrl{}
	err := db.Get(&result, "select * from tiny_url where shot_code = ?", shot)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
