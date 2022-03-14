package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/idzharbae/quickbid/src"
	"github.com/julienschmidt/httprouter"
)

type AttendanceService struct {
	attendanceUC src.AttendanceUC
}

func NewAttendanceService(attendanceUC src.AttendanceUC) *AttendanceService {
	return &AttendanceService{attendanceUC: attendanceUC}
}

type AttendanceRequest struct {
	Name string `json:"name"`
}

func (as *AttendanceService) GetHandles() []Handle {
	return []Handle{
		{
			Method: http.MethodPost,
			Path:   "/attendance",
			Handle: as.AttendanceHandler,
		},
	}
}

func (as *AttendanceService) AttendanceHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	var req AttendanceRequest
	err = json.Unmarshal([]byte(data), &req)
	if err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	err = as.attendanceUC.Attend(r.Context(), req.Name)
	if err != nil {
		log.Printf("[AttendanceHandler][attendanceUC] Attend: %s\n", err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Write([]byte("Success attendance!"))
}
