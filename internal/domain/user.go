package domain

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
}

type UserRepo interface {
	AddUser(username, email, password string) error
	GetByEmail(email string) (*User, error)
	CheckIfExists(user User) (bool, error)
}

type UserService struct {
	UserRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) AddUser(username, email, password string) error {
	return s.UserRepo.AddUser(username, email, password)
}

func (s *UserService) GetByEmail(email string) (*User, error) {
	return s.UserRepo.GetByEmail(email)
}

// func Login(UserId int64, password string) (*User, error) {
// 	return &User{}, nil
// }

// func Authorize(UserId int64) (token string) {
// 	token = "token"
// 	return token
// }

// func GetUsers() []User{

// }
