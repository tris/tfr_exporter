package main

// Response represents the top-level JSON response
type AirspaceResponse struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

// Data represents the data field in the JSON response
type Data struct {
	Color      Color      `json:"color"`
	Overview   Overview   `json:"overview"`
	Airports   []string   `json:"airports"`
	Classes    []string   `json:"classes"`
	Advisories []Advisory `json:"advisories"`
	Geometry   Geometry   `json:"geometry"`
	Region     string     `json:"region"`
}

// Color represents the color field in the Data struct
type Color struct {
	Name string `json:"name"`
	Hex  string `json:"hex"`
	RGB  []int  `json:"rgb"`
}

// Overview represents the overview field in the Data struct
type Overview struct {
	Short string `json:"short"`
	Full  string `json:"full"`
	Icon  string `json:"icon"`
}

// Advisory represents each item in the advisories array in the Data struct
type Advisory struct {
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Color       Color      `json:"color"`
	Description string     `json:"description"`
	Details     []Detail   `json:"details"`
	Geometry    string     `json:"geometry"` // This is a JSON string representing geometric data
	Distance    Distance   `json:"distance"`
	Properties  Properties `json:"properties"`
}

// Detail represents each item in the details array in the Advisory struct
type Detail struct {
	Type  string      `json:"type"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"` // Value can be string or int, hence using interface{}
}

// Distance represents the distance field in the Advisory struct
type Distance struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

// Properties represents the properties field in the Advisory struct
type Properties struct {
	ID           int    `json:"ID"`
	TFRID        string `json:"TFR_ID"`
	StartsAt     string `json:"STARTS_AT"`
	EndsAt       string `json:"ENDS_AT"`
	Reason       string `json:"REASON"`
	Link         string `json:"LINK"`
	Text         string `json:"TEXT"`
	DaysOfWeek   string `json:"DAYS_OF_WEEK"`
	Dist         string `json:"DIST"`
	ClosestPoint string `json:"CLOSEST_POINT"` // This is a JSON string representing a point
	TableName    string `json:"TABLE_NAME"`
}

// Geometry represents the geometry field in the Data struct
type Geometry struct {
	Format string `json:"format"`
	Data   string `json:"data"` // This is a JSON string representing geometric data
}
