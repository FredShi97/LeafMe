package api

import (
	"encoding/json"
	"github.com/2J/LeafMe/api/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetPlantByIdHandler(w http.ResponseWriter, r *http.Request) {
	plantId, _ := strconv.Atoi(urlParamAsString(r, "plantId"))

	plant := models.Plant{}

	err := plant.GetByID(plantId)

	if err != nil {
		writeErrorResponse(w, 404, "Plant not found")
		return
	}

	responseJSON, err := json.Marshal(plant)
	if err != nil {
		writeErrorResponse(w, 500, "125")
		return
	}
	writeJsonResponse(w, 200, responseJSON)
}

func UpdatePlantHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	plantID, _ := strconv.Atoi(urlParamAsString(r, "plantId"))

	type PlantUpdateForm struct {
		Name string `json:"name"`
	}
	plantUpdateForm := PlantUpdateForm{}
	json.Unmarshal(body, &plantUpdateForm)

	var plant models.Plant

	err := plant.GetByID(plantID)
	if err != nil {
		writeErrorResponse(w, 404, err.Error())
		return
	}

	plant.Name = plantUpdateForm.Name

	// Validate
	err = plant.Validate()
	if err != nil {
		writeErrorResponse(w, 400, err.Error())
		return
	}

	response := struct {
		Success bool `json:"success"`
	}{
		false,
	}
	responseJSON, _ := json.Marshal(response)

	err = models.UpdatePlantName(plantID, plant.Name)

	if err != nil {
		writeJsonResponse(w, 500, responseJSON)
		return
	}

	response.Success = true
	responseJSON, err = json.Marshal(response)
	writeJsonResponse(w, 200, responseJSON)
}
