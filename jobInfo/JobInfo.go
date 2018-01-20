package jobInfo

import (
	"database/sql"
	"fmt"
	"job_api/util"
	_ "github.com/lib/pq"

)

type JobInfo struct {
	Id      int
	City    string
	Date    string
	Content string
	Title   string
	JobType string
}

func QueryJobInfo(id string) JobInfo {
	db, err := sql.Open("postgres",
		"postgres://postgres:123456@10.16.6.94:5432/mutual_relation?sslmode=disable")

	util.CheckErr(err)

	//查询数据
	rows, err := db.Query("SELECT id, city, date, content, title, type  AS jobType FROM t_job_info WHERE id =" + id)
	util.CheckErr(err)
	jobInfo := JobInfo{}

	if rows.Next() {

		err = rows.Scan(&jobInfo.Id, &jobInfo.City, &jobInfo.Date, &jobInfo.Content, &jobInfo.JobType, &jobInfo.Title)
		util.CheckErr(err)

		fmt.Println(jobInfo)
	}

	db.Close()
	return jobInfo;
}


