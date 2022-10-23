package db

import (
	"time"

	sj "github.com/brianvoe/sjwt"
	"github.com/deta/deta-go/service/base"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
  Key string `json:"key"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
  Email string `json:"email"`
  Password string `json:"password"`
}

func GenUserToken(user User) (string) {
  info := &User {
    Key: user.Key,
    FirstName: user.FirstName,
    LastName: user.LastName,
    Email: user.Email,
  }
  claims, _ := sj.ToClaims(info)
  claims.SetExpiresAt(time.Now().Add(8760 * time.Hour))

  token := claims.Generate(JWTKey)
  return token
}

func ParseUserToken(token string) (User, error) {
  hasVerified := sj.Verify(token, JWTKey)

  if !hasVerified {
    return User {}, nil
  }

  claims, _ := sj.Parse(token)
  err := claims.Validate()
  user := User {}
  claims.ToStruct(&user)

  return user, err
}

func (user *User) Put() error {
  user.Key = GenKey()

  hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
  user.Password = string(hashedPassword)

  _, err := Users.Put(user)
  return err
}

func GetUser(q base.Query) (User, error) {
  var users []User

  _, err := Users.Fetch(&base.FetchInput{
    Q: q,
    Dest: &users,
    Limit: 1,
  })

  if (len(users) != 1) {
    return User {}, err
  }

  return users[0], err
}