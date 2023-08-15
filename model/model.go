package model

import "time"

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
