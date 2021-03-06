package routes

import (
	"errors"
	"yorha-api/datamodels"
	"yorha-api/services"

	"github.com/kataras/iris"
)

// Automatas returns list of the automatas. http://localhost:8080/automatas
func Automatas(service services.AutomataService) (results []datamodels.Automata) {
	return service.GetAll()
}

// AutomataByID return the current automata. http://localhost:8080/automatas/1
func AutomataByID(service services.AutomataService, id uint) (automata datamodels.Automata, found bool) {
	return service.GetByID(id)
}

// InsertAutomata create a new record in the database. http://localhost:8080/automatas
func InsertAutomata(ctx iris.Context, service services.AutomataService) (datamodels.Automata, error) {
	var automata datamodels.Automata

	if err := ctx.ReadJSON(&automata); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.Automata{}, errors.New("Json post request not correct at all")
	}

	return service.Create(automata.Name, automata.Occupation, automata.Photo, automata.RaceID)
}

// UpdateAutomata modify a current record in the database. http://localhost:8080/automatas/1
func UpdateAutomata(ctx iris.Context, service services.AutomataService, id uint) (datamodels.Automata, bool, error) {
	var automata datamodels.Automata

	if err := ctx.ReadJSON(&automata); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.Automata{}, true, errors.New("Json put request not correct at all")
	}

	return service.Update(id, automata.Name, automata.Occupation, automata.Photo, automata.RaceID)
}

// DeleteAutomata deletes the current automata. http://localhost:8080/automatas/1
func DeleteAutomata(service services.AutomataService, id uint) (automata datamodels.Automata, found bool) {
	return service.DeleteByID(id)
}
