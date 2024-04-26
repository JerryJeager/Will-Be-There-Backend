package manualwire

import (
	"github.com/JerryJeager/will-be-there-backend/config"
	"github.com/JerryJeager/will-be-there-backend/http"
	"github.com/JerryJeager/will-be-there-backend/service/invitees"
	"github.com/JerryJeager/will-be-there-backend/service/users"
	"github.com/JerryJeager/will-be-there-backend/service/event"
)

func GetUserRepository() *users.UserRepo {
	repo := config.GetSession()
	return users.NewUserRepo(repo)
}

func GetUserService(repo users.UserStore) *users.UserServ {
	return users.NewUserService(repo)
}

func GetUserController() *http.UserController {
	repo := GetUserRepository()
	service := GetUserService(repo)
	return http.NewUserController(service)
}
func GetEventRepository() *event.EventRepo {
	repo := config.GetSession()
	return event.NewEventRepo(repo)
}

func GetEventService(repo event.EventStore) *event.EventServ {
	return event.NewEventService(repo)
}

func GetEventController() *http.EventController {
	repo := GetEventRepository()
	service := GetEventService(repo)
	return http.NewEventController(service)
}
func GetInviteeRepository() *invitees.InviteeRepo {
	repo := config.GetSession()
	return invitees.NewInviteeRepo(repo)
}

func GetInviteeService(repo invitees.InviteeStore) *invitees.InviteeServ {
	return invitees.NewInviteeService(repo)
}

func GetInviteeController() *http.InviteeController {
	repo := GetInviteeRepository()
	service := GetInviteeService(repo)
	return http.NewInviteeController(service)
}
