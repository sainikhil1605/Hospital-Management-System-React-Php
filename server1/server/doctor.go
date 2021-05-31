package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type doctor struct{
	Id string `json:id`
	Name string `json:name`
	
	Department string `json:contact`
}
var doc []doctor
func GetDoctorEndPoint(w http.ResponseWriter, req *http.Request){
	rows,er:=db.Query("SELECT doctor_id,doc_name,department FROM `doctor`");
	doc:=[]doctor{};
	if(er==nil){
		fmt.Println("hi")
		var temp doctor
		for rows.Next(){
            rows.Scan(&temp.Id,&temp.Name,&temp.Department);
           
            doc=append(doc,temp);
         }
	}
	json.NewEncoder(w).Encode(doc);
}
func DeleteDoctorEndPoint(w http.ResponseWriter, req *http.Request){
	var p doctor
    err := json.NewDecoder(req.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	fmt.Println("Hi");
	fmt.Println(p);
	rows,er:=db.Query("DELETE FROM `doctor` where doctor_id="+p.Id);
	if(er==nil){
		fmt.Fprintf(w,"Success");
	}else{
		_=rows
		fmt.Fprintf(w,"Failed");
	}
    fmt.Fprintf(w, "Person: %+v", p)

}