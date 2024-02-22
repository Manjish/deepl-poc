package main

import (
    "deepl-poc/bootstrap"

    "github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()
    _ = bootstrap.RootApp.Execute()
}
