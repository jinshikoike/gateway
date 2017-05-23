package config

import "os"

// DBHost is db host
func DBHost() string {
	return os.Getenv("DB_HOST")
}

// DBPort db port
func DBPort() string {
	return os.Getenv("DB_PORT")
}

// DBUser db user to connect"
func DBUser() string {
	return os.Getenv("DB_USER")
}

// DBName is database name
func DBName() string {
	return os.Getenv("DB_NAME")
}

// DBPass return database password
func DBPass() string {
	return os.Getenv("DB_PASSWORD")
}
