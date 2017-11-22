package models

import(
	"regexp"
)

//user's model
type User struct {
	Username      string
	StudentId     string
	Phone         string
	Email         string
}

type ErrorMessage struct {
	Id            string
	Message       string
}

func IsAllValid(user User) []ErrorMessage{
	var errorMessageSlice []ErrorMessage
	if (isUsernameVaild(user.Username).Id != ""){
		errorMessageSlice = append(errorMessageSlice, isUsernameVaild(user.Username))
	}
	if (isStudentIdVaild(user.StudentId).Id != ""){
		errorMessageSlice = append(errorMessageSlice, isStudentIdVaild(user.StudentId))
	}
	if (isPhoneVaild(user.Phone).Id != ""){
		errorMessageSlice = append(errorMessageSlice, isPhoneVaild(user.Phone))
	}
	if (isEmailVaild(user.Email).Id != ""){
		errorMessageSlice = append(errorMessageSlice, isEmailVaild(user.Email))
	}
	return errorMessageSlice
}

func isUsernameVaild(username string) ErrorMessage{
	id := ""
	message := ""
	match, _ := regexp.MatchString(`[a-zA-Z]{1}\w{5,17}`, username)
	if !match {
		id = "errorname"
		message = "Attention:Username contains only a-z, A-Z or _, must begins with english letter, 6~8 characters"
	}
	return ErrorMessage{
		id, message,
	}
}

func isStudentIdVaild(stid string) ErrorMessage{
	id := ""
	message := ""
	match, _ := regexp.MatchString(`[^0]{1}\d{7}`, stid)
	if !match {
		id = "errorsid"
		message = "Attention:StudentId contains only 0~9, must begins without 0, 8 characters"
	}
	return ErrorMessage{
		id, message,
	}
}

func isPhoneVaild(phone string) ErrorMessage{
	id := ""
	message := ""
	match, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone)
	if !match {
		id = "errorphone"
		message = "Attention:Phone contains only 0~9, must begins without 0, 11 characters"
	}
	return ErrorMessage{
		id, message,
	}
}

func isEmailVaild(username string) ErrorMessage{
	id := ""
	message := ""
	match, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, username)
	if !match {
		id = "erroremail"
		message = "Attention:Email isn't allowed"
	}
	return ErrorMessage{
		id, message,
	}
}
