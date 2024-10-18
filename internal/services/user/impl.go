package user

func (u *UserService) Login() string {
	return u.repo.Login()
}

func (u *UserService) RetrieveData() string {
	return u.cache.RetrieveData()
}

func (u *UserService) Transfer() {

}
