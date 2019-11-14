package handlers

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"net/http"
)

type repositorySummary struct {
	Id			int //field needs to start with capital letters
	Address 	sql.NullString
	County 		sql.NullString
	District 	sql.NullString
	State 		sql.NullString
}

type repositories struct {
	Repositories []repositorySummary
}

func IndexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
				repos := repositories{}
				err := queryTrans(&repos, db)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			
				out, err := json.Marshal(repos)
				if err != nil{
					http.Error(w, err.Error(), 500)
					return
				}
				fmt.Fprintf(w, string(out))
			}
}

func queryTrans(repos *repositories, db *sql.DB) error {
	rows, err := db.Query(`
			SELECT
				id,
				address,
				county,
				district,
				state
			FROM transactions 
			LIMIT 10;`)
	
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		repo := repositorySummary{}
		err = rows.Scan(
			&repo.Id,
			&repo.Address,
			&repo.County,
			&repo.District,
			&repo.State,
		)
		if err != nil {
			return err
		}
		
		repos.Repositories = append(repos.Repositories, repo)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
