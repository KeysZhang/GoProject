package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {

	_, err := myegine.Insert(u)

	checkErr(err)

	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	everyone := make([]UserInfo,0)
	err := myegine.Find(&everyone)
	checkErr(err)
	return everyone
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	users := make([]UserInfo, 0)
	err := myegine.Id(id).Find(&users)
	checkErr(err)
	return &users[0]
}
