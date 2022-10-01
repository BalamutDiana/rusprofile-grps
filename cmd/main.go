package main

import (
	company "github.com/BalamutDiana/rusprofile-grps/internal/service"
)

func main() {
	inn := "2315214930"
	company.GetMainInfo(inn)
}
