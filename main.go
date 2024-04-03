package main

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"Artist"`
	Price string `json:"Price"`
}

var albums = []album{
	{ID: "1",Title: "Blue Train",}
}