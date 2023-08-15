package models

import "time"

type Productionschedule struct {
	Proddate         time.Time `json:"proddate"`
	Jobname          string    `json:"jobname"`
	Jobqueue         string    `json:"jobqueue"`
	Fullpathname     string    `json:"fullpathname"`
	Cardcount        int       `json:"cardcount"`
	Carriercount     int       `json:"carriercount"`
	Requestdate      time.Time `json:"requestdate"`
	Sentdate         time.Time `json:"sentdate"`
	Completedate     time.Time `json:"completedate"`
	Issues           int       `json:"issues"`
	Status           string    `json:"status"`
	Filesequence     int       `json:"filesequence"`
	Completedby      string    `json:"completedby"`
	Veridiedby       string    `json:"veridiedby"`
	Notes            string    `json:"notes"`
	Embossedcards    int       `json:"embossedcards"`
	Embossedcarriers int       `json:"embossedcarriers"`
	Heldcards        int       `json:"heldcards"`
	Heldcarriers     int       `json:"heldcarriers"`
}
