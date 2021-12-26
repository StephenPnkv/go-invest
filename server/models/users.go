package models

import (
  "fmt"
)

type User struct{
  Username string `bson:"username"`
  Password string `bson:"password"`
  email string `bson:"email"`
  phoneNumber string `bson:"phone_number"` 
}
