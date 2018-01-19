package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"job_api/jobInfo"
	"job_api/joblist"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/job/{id}/info", jobInfoHandler).Methods("GET")
	router.HandleFunc("/job/list", jobListHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func jobInfoHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	info := jobInfo.QueryJobInfo(id)
	log.Print(info.Id)
	json.NewEncoder(w).Encode(info)
}

func jobListHandler(w http.ResponseWriter, r *http.Request) {
	//city string,  key string,  limit int
	city := mux.Vars(r)["city"]
	key := mux.Vars(r)["key"]
	limit := mux.Vars(r)["limit"]

	joblistArray := joblist.QueryJobList(key, city, limit)
	json.NewEncoder(w).Encode(joblistArray)
}



