package user

type UserContext struct {
	UserService UserService
}

var Context UserContext

func init() {
	Context = UserContext{UserService: userCognitoService}
}
