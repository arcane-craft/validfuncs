package main

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	g "github.com/arcane-craft/validfuncs/generic"
	r "github.com/arcane-craft/validfuncs/records"
	c "github.com/arcane-craft/validfuncs/slices"
	s "github.com/arcane-craft/validfuncs/strings"
)

type UserInfo struct {
	Name   string   `json:"name"`
	Emails []string `json:"emails"`
	Phones []string `json:"phones,omitempty"`
}

type PassTip struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type RegisterRequest struct {
	Account  string     `json:"account"`
	Password string     `json:"password"`
	Detail   *UserInfo  `json:"detail,omitempty"`
	PassTips []*PassTip `json:"pass_tips,omitempty"`
}

var validation = validfuncs.MustCompile(
	g.Ref(g.Define(func(v RegisterRequest, R g.DefRet[RegisterRequest]) {
		R(
			r.JSONField(&v, &v.Account, s.RangeLen(5, 16), s.Charset(s.AlphaNum())),
			r.JSONField(&v, &v.Password, s.RangeLen(12, 18), s.Charset(s.PrintASCII())),
			r.JSONField(&v, &v.Detail, g.OmitNil(
				g.Ref(g.Define(func(v UserInfo, R g.DefRet[UserInfo]) {
					R(
						r.JSONField(&v, &v.Name, s.RangeLen(1, 20), s.Charset(s.AlphaNum())),
						r.JSONField(&v, &v.Emails, c.MinLenAnd(1, c.Each(s.Email(), s.MaxLen(60)))),
						r.JSONField(&v, &v.Phones, c.Each(s.MaxLen(15), s.Phone())),
					)
				})),
			)),
			r.JSONField(&v, &v.PassTips, c.Each(
				g.Ref(g.Define(func(v PassTip, R g.DefRet[PassTip]) {
					R(
						r.JSONField(&v, &v.Question, s.RangeLen(1, 50), s.Charset(s.PrintASCII(), s.Space())),
						r.JSONField(&v, &v.Answer, s.RangeLen(1, 50), s.Charset(s.PrintASCII(), s.Space())),
					)
				})),
			)),
		)
	})),
)

func main() {
	req := &RegisterRequest{
		Account:  "JinZhao",
		Password: "123456789111",
		Detail: &UserInfo{
			Name: "arcane-craft",
			Emails: []string{
				"arcane-craft@github.com",
			},
		},
		PassTips: []*PassTip{
			{
				Question: "where is your home?",
				Answer:   "你好啊",
			},
		},
	}
	if err := validation.Validate(req); err != nil {
		fmt.Println(err)
	}
}
