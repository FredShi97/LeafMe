package models

import (
	"database/sql"
	"errors"
	database "github.com/2J/LeafMe/api/db"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"time"
)

// LightingEvent https://github.com/2J/LeafMe/tree/master/api/models#lightingschedule
type LightingEvent struct {
	ID                 int       `json:"id" validate:"required"`
	PlantID            int       `json:"plant_id" validate:"required"`
	LightingScheduleID int       `json:"lighting_schedule_id" validate:"required"`
	StartTime          time.Time `json:"start_time"`
	Length             int       `json:"length" validate:"required"`
	Finished           bool      `json:"finished"`
}

// CreateLightingEvents creates batch of lighting events
func CreateLightingEvents(lightingEvents []LightingEvent) (err error) {
	db := database.Open()
	defer database.Close(db)

	query := "INSERT INTO lightingEvents (`plantId`, `lightingScheduleId`, `startTime`, `length`, `finished`) VALUES "
	vals := []interface{}{}
	for i, v := range lightingEvents {
		if i > 0 {
			query += ","
		}
		query += "(?,?,?,?,?)"

		vals = append(vals, v.PlantID, v.LightingScheduleID, v.StartTime, v.Length, v.Finished)
	}

	insForm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = insForm.Exec(vals...)

	return err
}

// CreateLightingEvent creates a lighting event
func (lightingSchedule *LightingSchedule) CreateLightingEvent() (err error) {
	// Create list of lighting events
	lightingEvents := []LightingEvent{}

	// Create lighting events until repeat ends
	t := *lightingSchedule.Schedule.Time
	for t.Before(*lightingSchedule.Schedule.RepeatEndDate) {
		lightingEvent := LightingEvent{}
		lightingEvent.PlantID = lightingSchedule.PlantID
		lightingEvent.LightingScheduleID = lightingSchedule.ID
		lightingEvent.StartTime = t
		lightingEvent.Length = lightingSchedule.Length
		lightingEvent.Finished = false

		lightingEvents = append(lightingEvents, lightingEvent)

		t = t.AddDate(0, 0, lightingSchedule.Schedule.RepeatDays)
	}

	// Chunk lighting events by 100 events each
	for i := 0; i < len(lightingEvents); i += 100 {
		chunk := lightingEvents[i:min(i+100, len(lightingEvents))]

		// Create events in database
		err = CreateLightingEvents(chunk)
		if err != nil {
			break
		}
	}

	return err
}

func (lightingEvent *LightingEvent) getRow(rows *sql.Rows) error {
	err := rows.Scan(
		&lightingEvent.ID,
		&lightingEvent.PlantID,
		&lightingEvent.LightingScheduleID,
		&lightingEvent.StartTime,
		&lightingEvent.Length,
		&lightingEvent.Finished,
	)

	return err
}

// GetLightingEventsByScheduleID gets all lighting events for schedule
func GetLightingEventsByScheduleID(scheduleID int) ([]LightingEvent, error) {
	lightingEvent := LightingEvent{}
	res := []LightingEvent{}

	db := database.Open()
	defer database.Close(db)
	rows, err := db.Query("SELECT * FROM lightingEvents WHERE lightingScheduleId = ? ORDER BY startTime ASC", scheduleID)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		lightingEvent.getRow(rows)

		if err != nil {
			return res, err
		}

		res = append(res, lightingEvent)
	}

	return res, nil
}

// GetLightingEventsByPlantID gets all events for plant
func GetLightingEventsByPlantID(plantID int) ([]LightingEvent, error) {
	lightingEvent := LightingEvent{}
	res := []LightingEvent{}

	db := database.Open()
	defer database.Close(db)
	rows, err := db.Query("SELECT * FROM lightingEvents WHERE plantId = ? ORDER BY startTime ASC", plantID)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		lightingEvent.getRow(rows)

		if err != nil {
			return res, err
		}

		res = append(res, lightingEvent)
	}

	return res, nil
}

// GetByID gets lighting event by ID
func (lightingEvent *LightingEvent) GetByID(id int) error {
	db := database.Open()
	defer database.Close(db)
	rows, err := db.Query("SELECT * FROM lightingEvents WHERE id = ?", id)
	if err != nil {
		return err
	}

	found := false
	for rows.Next() {
		found = true
		err = lightingEvent.getRow(rows)

		if err != nil {
			return err
		}
	}

	if !found {
		return errors.New("Event not found")
	}

	return nil
}

// DeleteLightingEventsByScheduleID deletes lighting events by schedule ID
func DeleteLightingEventsByScheduleID(scheduleID int) error {
	db := database.Open()
	defer database.Close(db)

	delForm, err := db.Prepare("DELETE FROM lightingEvents WHERE lightingScheduleId = ?")
	if err != nil {
		return err
	}

	_, err = delForm.Exec(
		scheduleID,
	)

	return err
}

// SetLightingEventFinished sets lighting event to finished
func SetLightingEventFinished(eventID int, finished bool) error {
	db := database.Open()
	defer database.Close(db)

	delForm, err := db.Prepare("UPDATE lightingEvents SET finished = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = delForm.Exec(
		finished,
		eventID,
	)

	return err
}
