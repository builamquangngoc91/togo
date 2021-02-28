package up

import (
	"github.com/manabie-com/togo/pkg/common/xerrors"
	"strings"
	"time"
)

type RegisterRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	MaxTodo  int    `json:"max_todo"`
}

func (r *RegisterRequest) Validate() error {
	if strings.TrimSpace(r.ID) == "" {
		return xerrors.ErrorM(xerrors.InvalidArgument, nil, "id is missing")
	}
	if strings.TrimSpace(r.Password) == "" {
		return xerrors.ErrorM(xerrors.InvalidArgument, nil, "password is missing")
	}
	if r.MaxTodo < 0 {
		return xerrors.ErrorM(xerrors.InvalidArgument, nil, "max_todo must be positive")
	}

	return nil
}

type RegisterResponse struct {
	ID      string `json:"id"`
	MaxTodo int    `json:"max_todo"`
}

type LoginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	if strings.TrimSpace(r.UserID) == "" {
		return xerrors.ErrorM(xerrors.InvalidArgument, nil, "user_id is missing")
	}
	if strings.TrimSpace(r.Password) == "" {
		return xerrors.ErrorM(xerrors.InvalidArgument, nil, "password is missing")
	}
	return nil
}

type LoginResponse string

type ListTasksRequest struct {
	CreatedDate string `json:"created_date"`
}

func (r *ListTasksRequest) Validate() error {
	_, err := time.Parse("2006-01-02", r.CreatedDate)
	if err != nil {
		return xerrors.ErrorM(xerrors.InvalidArgument, nil, "created_date doesn't match format yyyy-MM-dd")
	}
	return err
}

type ListTasksResponse []*Task

type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

type AddTaskRequest struct {
	Content string `json:"content"`
}

func (r *AddTaskRequest) Validate() error {
	if strings.TrimSpace(r.Content) == "" {
		return xerrors.ErrorM(xerrors.InvalidArgument, nil, "content is missing")
	}
	return nil
}
