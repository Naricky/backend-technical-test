package main

type StructData struct {
	Id        int    `json:Id`
	Name      string `json:Name`
	Address   string `json:Address`
	Latitude  int    `json:Latitude`
	Longitude int    `json:Longitude`
}

type IncomingRequest struct {
	Name    string `json:Name`
	Address string `json:Address`
}
