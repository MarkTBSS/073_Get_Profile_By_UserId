package databases

import (
	"log"

	_pkgConfig "github.com/MarkTBSS/073_Get_Profile_By_UserId/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func DbConnect(cfg _pkgConfig.IDbConfig) *sqlx.DB {
	// Connect
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to db failed: %v\n", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
