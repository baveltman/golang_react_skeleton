package main

type Context struct {
	CSRFToken 	string 	`json:"csrfToken"`
	Referer   	string 	`json:"referer"`
	Username	string	`json:"username"`
	Email 		string	`json:"email"`
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
	UserId        	string	`json:"userId" bson:"userId"`
	PhotoUrl	string	`json:"photoUrl"`
	IsAdmin		bool	`json:"isAdmin"`
}

type JsonContext struct {
	Json 			string
	IsProduction		bool
}