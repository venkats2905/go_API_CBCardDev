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

type Graphic struct {
	GraphicName         string    `json:"GRAPHICNAME"`
	GraphicRevisionNbr  int       `json:"GRAPHICREVISIONNBR"`
	GraphicDesc         string    `json:"GRAPHICDESC"`
	GraphicFileName     string    `json:"GRAPHICFILENAME"`
	GraphicImage        int64     `json:"GRAPHICIMAGE"`
	DefaultTopPos       float64   `json:"DEFAULTTOPPOS"`
	DefaultLeftPos      float64   `json:"DEFAULTLEFTPOS"`
	DefaultColorCode    string    `json:"DEFAULTCOLORCODE"`
	DefaultLocationCode string    `json:"DEFAULTLOCATIONCODE"`
	DefaultHeight       float64   `json:"DEFAULTHEIGHT"`
	DefaultWidth        float64   `json:"DEFAULTWIDTH"`
	EffectiveDate       time.Time `json:"EFFECTIVEDATE"`
	EndDate             time.Time `json:"ENDDATE"`
	LastUpdateTs        time.Time `json:"LASTUPDATETS"`
	LastUpdateUserId    string    `json:"LASTUPDATEUSERID"`
	VertOffset          int       `json:"VERTOOFFSET"`
	HorizOffset         int       `json:"HORIZOFFSET"`
}

type Platform struct {
	Idtype        string `json:"IDTYPE"`
	Priority      string `json:"PRIORITY"`
	Description   string `json:"DESCRIPTION"`
	Plasticstock  string `json:"PLASTICSTOCK"`
	Carrierstock  string `json:"CARRIERSTOCK"`
	Envelopestock string `json:"ENVELOPESTOCK"`
	Setup         string `json:"SETUP"`
	Fmodule1      string `json:"FMODULE1"`
	Fmodule2      string `json:"FMODULE2"`
	Fmodule3      string `json:"FMODULE3"`
	Bmodule1      string `json:"BMODULE1"`
	Bmodule2      string `json:"BMODULE2"`
	Jobgroup      string `json:"JOBGROUP"`
	Defaultqueue  string `json:"DEFAULTQUEUE"`
	Platformcode  string `json:"PLATFORMCODE"`
)
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
