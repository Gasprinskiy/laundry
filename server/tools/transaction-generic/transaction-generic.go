package transactiongeneric

import (
	"database/sql"
	"fmt"
	"laundry/internal/entity/global"

	"github.com/jmoiron/sqlx"
)

func HandleMethodWithTransaction[T any](
	db *sqlx.DB,
	fn func(tx *sqlx.Tx) (T, error),
	errMessage string,
) (T, error) {
	var result T

	tx, err := db.Beginx()
	if err != nil {
		return result, global.ErrInternalError
	}

	defer tx.Commit()

	result, err = fn(tx)
	if err != nil {
		if err == sql.ErrNoRows {
			err = global.ErrNoData
		} else {
			fmt.Printf("%s:%s", errMessage, err.Error())
			err = global.ErrInternalError
		}
		tx.Rollback()
	}

	return result, err
}
