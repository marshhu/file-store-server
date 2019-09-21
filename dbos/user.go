package dbos

import (
	"database/sql"
	"github.com/marshhu/file-store-server/util"
	"log"
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	UserNo     string    `json:"user_no"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	Avatar     string	 `json:"avatar"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	SignupAt   time.Time `json:"signup_at"`
	LastActive time.Time `json:"last_active"`
	Profile    string    `json:"profile"`
	Status     int       `json:"status"`
}

func AddUser(username string,password string,avatar string,phone string,email string,profile string) error{
	stmtIns, err := dbConn.Prepare("Insert into user(user_no,user_name,password,avatar,phone,email,signup_at,last_active,profile,status) values(?,?,?,?,?,?,?,?,?,?);")
	defer stmtIns.Close()
	if err != nil {
		return err
	}
	user := User{
		UserNo:util.NewUUID(),
		UserName:username,
		Password:password,
		Avatar:avatar,
		Phone:phone,
		Email:email,
		SignupAt:time.Now(),
		LastActive:time.Now(),
		Profile:profile,
		Status:1,
	}
	_,err = stmtIns.Exec(user.UserNo, user.UserName,user.Password,user.Avatar,user.Phone,user.Email,user.SignupAt,user.LastActive,user.Profile,user.Status)
	if err != nil{
		return  err
	}
	return nil
}

func CheckUser(username string,password string) (bool,error){
	stmtOut, err := dbConn.Prepare("select id from user where user_name = ? and password = ?;")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return false,err
	}
	var id int64
	err = stmtOut.QueryRow(username,password).Scan(&id)
	if err != nil && err != sql.ErrNoRows{
		return false,err
	}
	if err == sql.ErrNoRows || id <= 0 {
		return false,nil
	}
	return true,nil
}


func GetUserInfo(username string) (*User,error){
	stmtOut, err := dbConn.Prepare("select id,user_no,user_name,password,avatar,phone,email,signup_at,last_active,profile,status from user where user_name = ?;")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return nil,err
	}
	user := User{}
	var signupAt string
	var lastActive string
	err = stmtOut.QueryRow(username).Scan(&user.ID,&user.UserNo,&user.UserName,&user.Password,&user.Avatar,&user.Phone,&user.Email,&signupAt,&lastActive,&user.Profile,&user.Status)
	if err !=nil && err != sql.ErrNoRows{
		return nil, err
	}
	if err == sql.ErrNoRows{
		return nil,nil
	}
	user.SignupAt,_ = time.Parse("2006-01-02 15:04:05",signupAt)
	user.LastActive,_ = time.Parse("2006-01-02 15:04:05",lastActive)
	return &user,nil
}


