package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mcfiet/goDo/draw/model"
	drawService "github.com/mcfiet/goDo/draw/service"
	userService "github.com/mcfiet/goDo/user/service"
)

type DrawController struct {
	service     *drawService.DrawService
	userService *userService.UserService
}

func NewDrawController(service *drawService.DrawService, userService *userService.UserService) *DrawController {
	return &DrawController{service, userService}
}

func (controller *DrawController) GetAllDraws(w http.ResponseWriter, r *http.Request) {
	entries, err := controller.service.GetAllDraws()
	if err != nil {
		http.Error(w, "Fehler beim Holen der Daten", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(entries); err != nil {
		http.Error(w, "Fehler beim Codieren der Draws", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (controller *DrawController) CreateDraw(w http.ResponseWriter, r *http.Request) {
	var draw model.DrawResult

	userId := r.Context().Value("user_id").(string)

	draw.GiverId = uuid.MustParse(userId)

	fmt.Println("Draw.GiverId:", draw.GiverId)

	if err := controller.service.CreateDraw(&draw); err != nil {
		fmt.Println("Error1:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := controller.userService.FindById(draw.ReceiverId)
	if err != nil {
		fmt.Println("Error2:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, user.Username)
}

func (controller *DrawController) GetDrawByGiverId(w http.ResponseWriter, r *http.Request) {
	userId := uuid.MustParse(r.Context().Value("user_id").(string))
	fmt.Print(userId)

	entry, err := controller.service.GetDrawByGiverId(userId)
	if err != nil {
		fmt.Println("Error3:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := controller.userService.FindById(entry.ReceiverId)
	if err != nil {
		fmt.Println("Error4:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, user.Username)
}
