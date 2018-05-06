/********************************************************************
 * FileName:     user.go
 * Project:      BuddyApp
 * Module:       models
 * Company:      Budde
 * Developed by: Sourabh J Kalmady
 * Copyright and Disclaimer Notice Software:
 ********************************************************************/

package models

//USER registeration
type User struct {
	Id              int    //`db:"id"`
	Username      	string `db:"username"`      
	NickName 		string `db:"nick_name"` 	
	Password        string `db:"password"`
	EmailID         string `db:"email_id"`
	DOB         	string `db:"dob"`
	Gender         	string `db:"gender"`
	City         	string `db:"city"`
	Country         string `db:"country"`
	CollegeName     string `db:"college_name"`
	CompanyName     string `db:"company_name"`
	ProfilePic      string `db:"profile_pic"`
}

//Validate a light model before any operation on this
func InsertUser(user User) (err error) {
	

	err = Dbm.Insert(&user)

	return
}
