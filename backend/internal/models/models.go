package models

// Station представляет информацию о станции
type Station struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Train представляет информацию о поезде
type Train struct {
	Number          string `json:"number"`
	Number2         string `json:"number2"`
	Type            int    `json:"type"`
	TypeEx          int    `json:"typeEx"`
	Depth           int    `json:"depth"`
	New             bool   `json:"new"`
	ElReg           bool   `json:"elReg"`
	DeferredPayment bool   `json:"deferredPayment"`
	VarPrice        bool   `json:"varPrice"`
	Code0           int    `json:"code0"`
	Code1           int    `json:"code1"`
	BEntire         bool   `json:"bEntire"`
	TrainName       string `json:"trainName"`
	BFirm           bool   `json:"bFirm"`
	Brand           string `json:"brand"`
	Carrier         string `json:"carrier"`
	Route0          string `json:"route0"`
	Route1          string `json:"route1"`
	RouteCode0      int    `json:"routeCode0"`
	RouteCode1      int    `json:"routeCode1"`
	TrDate0         string `json:"trDate0"`
	TrTime0         string `json:"trTime0"`
	Station0        string `json:"station0"`
	Station1        string `json:"station1"`
	Date0           string `json:"date0"`
	Time0           string `json:"time0"`
	Date1           string `json:"date1"`
	Time1           string `json:"time1"`
	TimeInWay       string `json:"timeInWay"`
	FlMsk           int    `json:"flMsk"`
	TrainID         int    `json:"train_id"`
	Cars            []Car  `json:"cars"`
	CarNumeration   string `json:"carNumeration"`
	AddCompLuggage  bool   `json:"addCompLuggage"`
	AddHandLuggage  bool   `json:"addHandLuggage"`
}

// Car представляет информацию о вагоне
type Car struct {
	CarDataType int     `json:"carDataType"`
	IType       int     `json:"itype"`
	Type        string  `json:"type"`
	TypeLoc     string  `json:"typeLoc"`
	FreeSeats   int     `json:"freeSeats"`
	PT          int     `json:"pt"`
	Tariff      float64 `json:"tariff"`
	ServCls     string  `json:"servCls"`
}

// SearchRequest представляет параметры поиска поездов
type SearchRequest struct {
	FromCode string `json:"fromCode" binding:"required"`
	ToCode   string `json:"toCode" binding:"required"`
	Date     string `json:"date" binding:"required"`
}

// SearchResponse представляет результат поиска поездов
type SearchResponse struct {
	Trains []Train `json:"trains"`
}

// ErrorResponse представляет структуру ответа с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}

// HealthResponse представляет структуру ответа для проверки здоровья сервиса
type HealthResponse struct {
	Status string `json:"status"`
}
