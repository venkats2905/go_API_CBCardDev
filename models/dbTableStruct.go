package models

import "time"

type Productioncard struct {
	Contractnbr      string    `json:"contractnbr"`
	Requestdate      time.Time `json:"requestdate"`
	Status           string    `json:"status"`
	Statusdate       time.Time `json:"statusdate"`
	Statusby         string    `json:"statusby"`
	Searchcode       uint      `json:"searchcode"`
	Cardcount        uint      `json:"cardcount"`
	Jobname          string    `json:"jobname"`
	Producedby       string    `json:"producedby"`
	Produceddate     time.Time `json:"produceddate"`
	Scheduleddate    time.Time `json:"scheduleddate"`
	Cardtemplatecode string    `json:"cardtemplatecode"`
	Groupnbr         string    `json:"groupnbr"`
	Suffixnbr        string    `json:"suffixnbr"`
	Matrldist        string    `json:"matrldist"`
	Trancd           string    `json:"trancd"`
	Reasoncd         string    `json:"reasoncd"`
	Reptype          string    `json:"reptype"`
	Litcode          string    `json:"litcode"`
}
