package shortener

import "database/sql"

type Url struct {
	ID         int64  `json:"id"`
	URL        string `json:"url"`
	SHORT_ID   string `json:"shortId"`
	CREATED_AT string `json:"createdAt"`
	UPDATED_AT string `json:"updatedAt"`
	DELETED_AT string `json:"deletedAt"`
}

func (url *Url) findById(db *sql.DB) error {
	return db.QueryRow("SELECT id, url, shortid, created_at, updated_at, COALESCE(deleted_at, '') as deleted_at FROM shortened_urls WHERE id = ?", url.ID).Scan(&url.ID, &url.URL, &url.SHORT_ID, &url.CREATED_AT, &url.UPDATED_AT, &url.DELETED_AT)
}

func (url *Url) findByUrl(db *sql.DB) error {
	return db.QueryRow("SELECT id, url, shortid, created_at, updated_at, COALESCE(deleted_at, '') as deleted_at FROM shortened_urls WHERE url = ?", url.URL).Scan(&url.ID, &url.URL, &url.SHORT_ID, &url.CREATED_AT, &url.UPDATED_AT, &url.DELETED_AT)
}

func (url *Url) findByShortid(db *sql.DB) error {
	return db.QueryRow("SELECT id, url, shortid, created_at, updated_at, COALESCE(deleted_at, '') as deleted_at FROM shortened_urls WHERE shortid = ?", url.SHORT_ID).Scan(&url.ID, &url.URL, &url.SHORT_ID, &url.CREATED_AT, &url.UPDATED_AT, &url.DELETED_AT)
}

func (url *Url) create(db *sql.DB) (sql.Result, error) {
	return db.Exec(`INSERT INTO shortened_urls (url, shortid, created_at, updated_at) VALUES(?, ?, now(), now())`, url.URL, url.SHORT_ID)
}
