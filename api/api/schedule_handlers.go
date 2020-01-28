package api

import (
	"encoding/json"
	"github.com/2J/LeafMe/api/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetLightingScheduleByIdHandler(w http.ResponseWriter, r *http.Request) {
	//plantId, _ := strconv.Atoi(urlParamAsString(r, "plantId"))
	scheduleId, _ := strconv.Atoi(urlParamAsString(r, "scheduleId"))

	lightingSchedule := models.LightingSchedule{}
	err := lightingSchedule.GetById(scheduleId)

	if err != nil {
		writeErrorResponse(w, 500, "Failed to get: "+err.Error())
		return
	}

	responseJson, err := json.Marshal(lightingSchedule)
	if err != nil {
		writeErrorResponse(w, 500, "125")
		return
	}
	writeJsonResponse(w, 200, responseJson)
}

func GetWateringScheduleByIdHandler(w http.ResponseWriter, r *http.Request) {
	//plantId, _ := strconv.Atoi(urlParamAsString(r, "plantId"))
	scheduleId, _ := strconv.Atoi(urlParamAsString(r, "scheduleId"))

	wateringSchedule := models.WateringSchedule{}
	err := wateringSchedule.GetById(scheduleId)

	if err != nil {
		writeErrorResponse(w, 500, "Failed to get: "+err.Error())
		return
	}

	responseJson, err := json.Marshal(wateringSchedule)
	if err != nil {
		writeErrorResponse(w, 500, "125")
		return
	}
	writeJsonResponse(w, 200, responseJson)
}

func GetSchedulesByPlantIdHandler(w http.ResponseWriter, r *http.Request) {
	plantId, _ := strconv.Atoi(urlParamAsString(r, "plantId"))

	lightingSchedules, err := models.GetLightingSchedulesByPlantId(plantId)
	if err != nil {
		panic(err.Error())
	}
	wateringSchedules, err := models.GetWateringSchedulesByPlantId(plantId)
	if err != nil {
		panic(err.Error())
	}

	response := struct {
		LightingSchedules []models.LightingSchedule `json:"lighting_schedules"`
		WateringSchedules []models.WateringSchedule `json:"watering_schedules"`
	}{
		lightingSchedules,
		wateringSchedules,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		writeErrorResponse(w, 500, "125")
		return
	}
	writeJsonResponse(w, 200, responseJson)
}

func CreateLightingScheduleHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	plantID, _ := strconv.Atoi(urlParamAsString(r, "plantId"))

	lightingSchedule := models.LightingSchedule{}
	json.Unmarshal(body, &lightingSchedule)

	lightingSchedule.PlantID = plantID

	// TODO: Validate

	response := struct {
		Success bool `json:"success"`
		Id      int  `json:"id"`
	}{
		false,
		0,
	}
	responseJson, _ := json.Marshal(response)

	var err error
	response.Id, err = lightingSchedule.Create()

	if err != nil {
		writeJsonResponse(w, 500, responseJson)
		return
	}

	response.Success = true
	responseJson, err = json.Marshal(response)
	writeJsonResponse(w, 200, responseJson)
}

func DeleteLightingScheduleHandler(w http.ResponseWriter, r *http.Request) {
	scheduleID, _ := strconv.Atoi(urlParamAsString(r, "scheduleId"))

	err := models.DeleteLightingSchedule(scheduleID)

	if err != nil {
		writeErrorResponse(w, 500, "Failed delete: "+err.Error())
		return
	}

	writeJsonResponse(w, 200, []byte(""))
}

func CreateWateringScheduleHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	plantID, _ := strconv.Atoi(urlParamAsString(r, "plantId"))

	wateringSchedule := models.WateringSchedule{}
	json.Unmarshal(body, &wateringSchedule)

	wateringSchedule.PlantID = plantID

	// TODO: Validate

	response := struct {
		Success bool `json:"success"`
		Id      int  `json:"id"`
	}{
		false,
		0,
	}
	responseJson, _ := json.Marshal(response)

	var err error
	response.Id, err = wateringSchedule.Create()

	if err != nil {
		writeJsonResponse(w, 500, responseJson)
		return
	}

	response.Success = true
	responseJson, err = json.Marshal(response)
	writeJsonResponse(w, 200, responseJson)
}

func DeleteWateringScheduleHandler(w http.ResponseWriter, r *http.Request) {
	scheduleID, _ := strconv.Atoi(urlParamAsString(r, "scheduleId"))

	err := models.DeleteWateringSchedule(scheduleID)

	if err != nil {
		writeErrorResponse(w, 500, "Failed delete: "+err.Error())
		return
	}

	writeJsonResponse(w, 200, []byte(""))
}
