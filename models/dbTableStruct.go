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

type ProcessedFile struct {
	Fileid          int       `json:"fielid"`
	Source_system   string    `json:"source_system"`
	Platform        string    `json:"platform"`
	Isproduction    string    `json:"isproduction"`
	Header_date     time.Time `json:"header_date"`
	Process_date    time.Time `json:"process_date"`
	Start_seq       int       `json:"start_seq"`
	End_seq         int       `json:"end_seq"`
	File_name       string    `json:"file_name"`
	Mabx_csm_fileid string    `json:"mabx_csm_fileid"`
}
