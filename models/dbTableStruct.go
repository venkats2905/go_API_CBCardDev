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

type Trackingsequence_archive struct {
	SID              int       `json:"SID" bun:"SID"`
	CONTRACTNBR      string    `json:"CONTRACTNBR" bun:"CONTRACTNBR"`
	SUBSEQ           string    `json:"SUBSEQ" bun:"SUBSEQ"`
	GROUPSEQ         string    `json:"GROUPSEQ" bun:"GROUPSEQ"`
	GROUPNBR         string    `json:"GROUPNBR" bun:"GROUPNBR"`
	SUFFIXNBR        string    `json:"SUFFIXNBR" bun:"SUFFIXNBR"`
	HOMEPLAN         string    `json:"HOMEPLAN" bun:"HOMEPLAN"`
	IMB              string    `json:"IMB" bun:"IMB"`
	MATRLDIST        string    `json:"MATRLDIST" bun:"MATRLDIST"`
	DISPCODE         string    `json:"DISPCODE" bun:"DISPCODE"`
	MAILDATE         time.Time `json:"MAILDATE" bun:"MAILDATE"`
	SENTFLAG         string    `json:"SENTFLAG" bun:"SENTFLAG"`
	CARDTEMPLATECODE string    `json:"CARDTEMPLATECODE" bun:"CARDTEMPLATECODE"`
	SETUPNAME        string    `json:"SETUPNAME" bun:"SETUPNAME"`
	REQ_DATE         time.Time `json:"REQ_DATE" bun:"REQ_DATE"`
	SEARCHCODE       int       `json:"SEARCHCODE" bun:"SEARCHCODE"`
	RELEASE_DT       time.Time `json:"RELEASE_DT" bun:"RELEASE_DT"`
	REPTTYPE         string    `json:"REPTTYPE" bun:"REPTTYPE"`
	FEPFLAG          string    `json:"FEPFLAG" bun:"FEPFLAG"`
}
