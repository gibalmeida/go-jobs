// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type NewAdmin struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewApplicant struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewManager struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	DepartmentID string `json:"departmentId"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Role string

const (
	RoleAdmin     Role = "ADMIN"
	RoleManager   Role = "MANAGER"
	RoleApplicant Role = "APPLICANT"
)

var AllRole = []Role{
	RoleAdmin,
	RoleManager,
	RoleApplicant,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleManager, RoleApplicant:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}