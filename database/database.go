package database

var connection string

// function init berfungsi untuk menjalan perintah-perintah yang ada di dalamnya ketika package itu di panggil dari package lain
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
