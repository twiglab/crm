// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type PolyDetailReq struct {
	PolyCode string `json:"polyCode"`
}

type PolyDetailResp struct {
	PolyCode  string    `json:"polyCode"`
	MallCode  string    `json:"mallCode"`
	Title     string    `json:"title"`
	Name      string    `json:"name"`
	Memo      string    `json:"memo"`
	LogoPic   string    `json:"logoPic"`
	Pic1      string    `json:"pic1"`
	Pic2      string    `json:"pic2"`
	Pic3      string    `json:"pic3"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Type      int32     `json:"type"`
}

type PolyPageReq struct {
	MailCode   string  `json:"mailCode"`
	MemberCode string  `json:"memberCode"`
	Last       *string `json:"last,omitempty"`
	Limit      int32   `json:"limit"`
}

type PolyTitlelResp struct {
	PolyCode  string    `json:"polyCode"`
	Title     string    `json:"title"`
	Name      string    `json:"name"`
	LogoPic   string    `json:"logoPic"`
	Pic1      string    `json:"pic1"`
	Pic2      string    `json:"pic2"`
	Pic3      string    `json:"pic3"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Type      int32     `json:"type"`
}

type QualificationReq struct {
	MemberCode string `json:"memberCode"`
	PolyCode   string `json:"polyCode"`
}

type QualificationResp struct {
	MemberCode string `json:"memberCode"`
	PolyCode   string `json:"polyCode"`
	Token      string `json:"token"`
	Result     int32  `json:"result"`
}

type Query struct {
}