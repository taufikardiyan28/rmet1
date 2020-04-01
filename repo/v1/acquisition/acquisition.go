package AcquisitionRepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	db "github.com/taufikardiyan28/rmet1/db"
	h "github.com/taufikardiyan28/rmet1/helper"
	AuditModel "github.com/taufikardiyan28/rmet1/model/audit"
	BuildingModel "github.com/taufikardiyan28/rmet1/model/building"
	BuildStaffModel "github.com/taufikardiyan28/rmet1/model/buildstaff"
	DocumentModel "github.com/taufikardiyan28/rmet1/model/document"
	OwnerModel "github.com/taufikardiyan28/rmet1/model/owner"
	RoomTypeModel "github.com/taufikardiyan28/rmet1/model/roomtype"
)

type (
	Repo struct {
		db.DB
	}

	BuildingData struct {
		BuildingModel.Building
		OwnerModel.Owner
		BuildStaffModel.BuildStaff
		AuditStatus string `db:"audit_status" json:"audit_status"`
		Room_Count  int    `db:"room_count" json"room_count"`
	}

	SummaryData struct {
		AuditStatus string `db:"audit_status" json:"audit_status"`
		Total       int    `db:"total" json:"total"`
	}

	PlaceIDResponse struct {
		Candidates []struct {
			PlaceID string `json:"place_id"`
		} `json:"candidates"`
		ErrorMessage string `json:"error_message"`
		Status       string `json:"status"`
	}

	PlaceIDDetailResponse struct {
		HTMLAttributions []interface{} `json:"html_attributions"`
		Result           struct {
			AddressComponents []struct {
				LongName  string   `json:"long_name"`
				ShortName string   `json:"short_name"`
				Types     []string `json:"types"`
			} `json:"address_components"`
			AdrAddress       string `json:"adr_address"`
			FormattedAddress string `json:"formatted_address"`
			Geometry         struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				Viewport struct {
					Northeast struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"northeast"`
					Southwest struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"southwest"`
				} `json:"viewport"`
			} `json:"geometry"`
			Icon   string `json:"icon"`
			ID     string `json:"id"`
			Name   string `json:"name"`
			Photos []struct {
				Height           int      `json:"height"`
				HTMLAttributions []string `json:"html_attributions"`
				PhotoReference   string   `json:"photo_reference"`
				Width            int      `json:"width"`
			} `json:"photos"`
			PlaceID   string   `json:"place_id"`
			Reference string   `json:"reference"`
			Scope     string   `json:"scope"`
			Types     []string `json:"types"`
			URL       string   `json:"url"`
			UtcOffset int      `json:"utc_offset"`
			Vicinity  string   `json:"vicinity"`
			Website   string   `json:"website"`
		} `json:"result"`
		ErrorMessage string `json:"error_message"`
		Status       string `json:"status"`
	}
)

func (r *Repo) getDefaultQuery() string {
	return `SELECT * FROM (SELECT *, CASE  WHEN build_audit=0 THEN 'Belum Audit' WHEN build_total_room>build_audit THEN 
			'Audit Ulang' ELSE 'Completed' END AS audit_status FROM building) as t `
}

func (r *Repo) getSummaryQuery() string {
	return `SELECT audit_status, COUNT(*) AS total FROM (SELECT *, CASE  WHEN build_audit=0 THEN 'Belum Audit' WHEN build_total_room>build_audit THEN 
			'Audit Ulang' ELSE 'Completed' END AS audit_status FROM building) AS t
			GROUP BY audit_status
			UNION 
			SELECT 'All' AS audit_status, COUNT(*) FROM building`
}

func (r *Repo) List(paging *h.PagingData) ([]BuildingData, int, error) {
	strSql := r.getDefaultQuery()

	data := []BuildingData{}
	period := h.PeriodFilter{
		DateColumn: "build_create_date",
		StartDate:  paging.StartCreateDate,
		EndDate:    paging.EndCreateDate + " 23:59:59",
	}
	totalData, err := r.GetPaging(&data, strSql, paging, period)
	return data, totalData, err
}

func (r *Repo) Summary() ([]SummaryData, error) {
	strSql := r.getSummaryQuery()

	data := []SummaryData{}

	err := r.Select(&data, strSql)
	return data, err
}

func (r *Repo) GetDetail(buildId int) ([]BuildingData, error) {
	strSql := `SELECT a.*, CASE  WHEN build_audit=0 THEN 'Belum Audit' WHEN build_total_room>build_audit THEN 
	'Audit Ulang' ELSE 'Completed' END AS audit_status,  
	room_count, o.owner_name, o.owner_phone, s.build_staff_name, s.build_staff_phone FROM building a LEFT JOIN 
	(SELECT roomtype_build_id,  COUNT(*) AS room_count FROM room_type 
	GROUP BY roomtype_build_id) AS b ON a.build_id=b.roomtype_build_id
	LEFT JOIN ` + "`owner`" + ` o ON a.build_id=o.owner_building_id  AND o.owner_del_status = '0'
	LEFT JOIN build_staff s ON a.build_id = s.build_staff_build_id AND s.build_staff_del_status = '0'
	WHERE a.build_id=?;`

	data := []BuildingData{}

	err := r.Select(&data, strSql, buildId)
	return data, err
}

func (r *Repo) GetDocument(buildId int) ([]DocumentModel.Document, error) {
	strSql := `SELECT * FROM document WHERE doc_build_id=? AND doc_del_status='0'`

	data := []DocumentModel.Document{}

	err := r.Select(&data, strSql, buildId)
	return data, err
}

func (r *Repo) GetRoomType(buildId int) ([]RoomTypeModel.RoomType, error) {
	strSql := `SELECT * FROM room_type WHERE roomtype_build_id=? AND roomtype_del_status='0'`

	data := []RoomTypeModel.RoomType{}

	err := r.Select(&data, strSql, buildId)
	return data, err
}

func (r *Repo) GetAuditLog(buildId int) ([]AuditModel.Audit, error) {
	strSql := `SELECT * FROM audit WHERE audit_build_id=? AND audit_del_status='0'`

	data := []AuditModel.Audit{}

	err := r.Select(&data, strSql, buildId)
	return data, err
}

func (r *Repo) GetGooglePlaceId(param string) (string, error) {
	url := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?key=%s&input=%s&inputtype=textquery&language=id"
	url = fmt.Sprintf(url, h.Config.GoogleApiKey, param)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res := PlaceIDResponse{}
	err = json.Unmarshal(body, &res)
	placeId := ""
	if res.Status == "OK" {
		placeId = res.Candidates[0].PlaceID
	} else {
		err = errors.New(res.ErrorMessage)
	}
	return placeId, err
}

func (r *Repo) GetLatLng(param string) (string, string, error) {
	url := "https://maps.googleapis.com/maps/api/place/details/json?placeid=%s&key=%s"
	url = fmt.Sprintf(url, param, h.Config.GoogleApiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	res := PlaceIDDetailResponse{}
	err = json.Unmarshal(body, &res)
	lat, lng := "", ""
	if res.Status == "OK" {
		lat = fmt.Sprintf("%f", res.Result.Geometry.Location.Lat)
		lng = fmt.Sprintf("%f", res.Result.Geometry.Location.Lng)
	} else {
		fmt.Println("error", res.ErrorMessage)
		err = errors.New(res.ErrorMessage)
	}
	return lat, lng, err
}
