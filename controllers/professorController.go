package controllers

import (
	"ProjetoFinal/models"
	"fmt"
	"net/http"
)

func GetProfessor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	gabriel := models.Professor{
		Name:  "Gabriel",
		Email: "gabriel@gmail.com",
		CPF:   "12312312311",
	}

	fmt.Fprintf(w, "%+v", gabriel)
}

func CreateProfessor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Criado")
}

func UpdateProfessor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Atualizado")
}

func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Deletado")
}
