## Installation

#### Clone the source

```bash
git clone https://github.com/taufikardiyan28/rmet1.git
```

#### Setup dependencies

```bash
go build
```

#### Configuration
Edit config.yaml file and change your port and mysql database configuration
and don't forget to create your database with dump.sql

#### Run the app
```
go run main.go
or
./rmet1
```

### Usage


* List All Acquisition
  path: /v1/form/acquisition/list \
  method: POST \
  Example Post Data: \
  {
    "show":100,
    "page":1,
    "search":"[{\"column\":\"audit_status\", \"op\":\"eq\",\"value\":\"Completed\"}]",
    "start_create_date":"",
    "end_create_date": ""
  }
  
 * Acquisition Detail
  path: /v1/form/acquisition/detail \
  method: POST \
  Example Post Data: \
  {
    "build_id": 509
  }
  
* List All Acquisition
  path: /v1/form/acquisition/list \
  method: POST \
  Example Post Data: \
  {
    "show":100,
    "page":1,
    "search":"[{\"column\":\"audit_status\", \"op\":\"eq\",\"value\":\"Completed\"}]",
    "start_create_date":"",
    "end_create_date": ""
  }
  
 * User List
  path: /v1/user/list \
  method: POST \
  Example Post Data: \
  {
    "show":100,
    "page":1,
    "search":"",
    "start_create_date":"",
    "end_create_date": ""
  }
  
  * Find Google Place ID
  path: /v1/form/acquisition/find-place-id \
  method: POST \
  Example Post Data: \
  {
    "search": "Jakarta"
  }
  
  * Get Langitude and Longitude from Google Place ID
  path: /v1/form/acquisition/get-latlng \
  method: POST \
  Example Post Data: \
  {
    "place_id": "ChIJnUvjRenzaS4RoobX2g-_cVMddd"
  }
  
  * Download Document
  path: /get/file \
  method: POST \
  Example Post Data: \
  {
    "file_name": "documents_2_2020-03-18_02-42-38_144-Article Text-298-2-10-20190510.pdf"
  }
  
  * Room Type List
  path: /v1/roomtype/list \
  method: POST \
  Example Post Data: \
  {
    "show":5,
    "page":1,
    "search":"",
    "start_create_date":"",
    "end_create_date": ""
  }
  
  
  * Set Up Audit (Change RoomType)
  path: /v1/form/acquisition/updateroom \
  method: PUT \
  Example Post Data: \
  {
    "build_id": 521,
    "data": [
      {"roomtype_id" : 358, "roomtype_name": "Comfort"},
      {"roomtype_id" : 0, "roomtype_name": "Comfort King Bed"}
    ]
  }
 
