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
