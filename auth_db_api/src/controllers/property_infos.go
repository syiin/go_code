package controllers

import (
	"encoding/json"
	"models"
	"net/http"
)

func FetchPropertyInfos(w http.ResponseWriter, r *http.Request) {
	var propertyinfos []models.PropertyInfo
	db.Table("property_infos").Limit(10).Find(&propertyinfos)
	json.NewEncoder(w).Encode(propertyinfos)
}
