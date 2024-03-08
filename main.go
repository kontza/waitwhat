/*
Copyright © 2023 Juha Ruotsalainen <juha.ruotsalainen@iki.fi>
*/
package main

import (
	"embed"
	"waitwhat/cmd"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	cmd.Execute(assets)
}
