package joblist

import (
	"strconv"
	"database/sql"
	"log"
	"job_api/util"
	_ "github.com/lib/pq"

)

type JobList struct {
	Id      int
	City    string
	Date    string
	Title   string
	JobType string
}

func QueryJobList(city string, key string, limit string) [10]JobList {

	querySQl := "SELECT id, city, date,  title, type  AS jobType FROM t_job_info where 1 = 1"

	if city != "" {
		querySQl += (" and city = " + city)
	}

	if key != "" {
		querySQl += (" and title like %" + city + "%")
	}

	if limit == "" {
		limit = "0"
	}
	querySQl += (" limit " + strconv.Itoa(10) + " OFFSET " + limit)

	log.Print(querySQl)

	db, err := sql.Open("postgres",
		"postgres://postgres:123456@10.16.6.94:5432/mutual_relation?sslmode=disable")

	util.CheckErr(err)

	//查询数据
	rows, err := db.Query(querySQl)
	util.CheckErr(err)
	jobList := JobList{}

	listArray := [10]JobList{}
	i := 0
	for rows.Next() {

		err = rows.Scan(&jobList.Id, &jobList.City, &jobList.Date, &jobList.Title, &jobList.JobType)
		util.CheckErr(err)
		listArray[i] = jobList
		i += 1
	}

	db.Close()
	return listArray;
}
