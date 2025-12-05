package user

// func GetModels(helper QueryHelper) []*Entity {
// 	return entities.GetModels[*Entity](helper)
// }

// func GetModel(helper QueryHelper) *Entity {
// 	return entities.GetModel[*Entity](helper)
// }

// func GetCount(helper QueryHelper) int64 {
// 	return entities.GetCount[*Entity](helper)
// }

// func Save(model *Entity) error {
// 	return entities.Save(model)
// }

// func Delete(model *Entity) error {
// 	return entities.Delete(model)
// }

// func GetModelByID(id int64, helper QueryHelper) *Entity {
// 	helper.ID = ID(id)
// 	return GetModel(helper)
// }

// func GetCountOfFinishedEvents(eventID, userID int64) int64 {
// 	query := efw.DBTransaction[*Entity]()
// 	exec := query.Raw(
// 		"SELECT COUNT(*) AS cnt FROM log_events WHERE event_id = ? AND user_id = ? AND is_completed = TRUE",
// 		eventID, userID,
// 	)

// 	var result struct {
// 		Cnt int64 `gorm:"column:cnt"`
// 	}

// 	if err := exec.Scan(context.Background(), &result); err != nil {
// 		if !errors.Is(err, gorm.ErrRecordNotFound) {
// 			efwlog.LogError(err)
// 		}
// 		return 0
// 	}
// 	return result.Cnt
// }

// func GetCountOfFinishedEventsByGroupID(groupID, userID int64) int64 {
// 	query := efw.DBTransaction[*Entity]()
// 	exec := query.Raw(
// 		`SELECT COUNT(*) AS cnt
// 		 FROM log_events
// 		 JOIN events ON events.id = log_events.event_id
// 		 WHERE events.group_id = ? AND log_events.user_id = ? AND log_events.is_completed = TRUE`,
// 		groupID, userID,
// 	)

// 	var result struct {
// 		Cnt int64 `gorm:"column:cnt"`
// 	}

// 	if err := exec.Scan(context.Background(), &result); err != nil {
// 		if !errors.Is(err, gorm.ErrRecordNotFound) {
// 			efwlog.LogError(err)
// 		}
// 		return 0
// 	}
// 	return result.Cnt
// }

// func GetLastFinishedEventByGroupID(groupID, userID int64) *Entity {
// 	query := efw.DBTransaction[*Entity]()
// 	exec := query.Raw(
// 		`SELECT log_events.*
// 		 FROM log_events
// 		 JOIN events ON events.id = log_events.event_id
// 		 WHERE events.group_id = ? AND log_events.user_id = ? AND log_events.is_completed = TRUE
// 		 ORDER BY log_events.id DESC
// 		 LIMIT 1`,
// 		groupID, userID,
// 	)

// 	entity := &Entity{}
// 	if err := exec.Scan(context.Background(), entity); err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return &Entity{}
// 		}
// 		efwlog.LogError(err)
// 		return &Entity{}
// 	}
// 	return entity
// }
