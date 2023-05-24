package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	api "github.com/seekr-osint/seekr/api"
	"github.com/seekr-osint/seekr/api/discord"
	"github.com/seekr-osint/seekr/api/server"
	"github.com/seekr-osint/seekr/api/webserver"
	"github.com/seekr-osint/seekr/seekrplugin"
)

// Web server content
//
//go:embed web/*
var content embed.FS

var dataBase = make(api.DataBase)

var version string

func main() {
	err := discord.Rich()
	if err != nil {
		fmt.Printf("%s", err)
	}
	if version != "" {
		fmt.Printf("Welcome to seekr v%s\n", version)
	} else {
		fmt.Printf("Welcome to seekr unstable\nplease note that this version of seekr is NOT officially supported\n")
	}
	// dir := flag.String("dir", "./web", "dir where the html source code is located")
	ip := flag.String("ip", "localhost", "Ip to serve api + webServer on (0.0.0.0 or localhost usually)")
	data := flag.String("db", "data", "Database location")
	port := flag.Uint64("port", 8569, "Port to serve API on")
	enableWebserver := flag.Bool("webserver", true, "Enable the webserver")

	forcePort := flag.Bool("forcePort", false, "forcePort")
	//enableWebserver := flag.Bool("webserver", true, "Enable the webserver")
	enableApiServer := true
	// webserverPort := flag.String("webserverPort", "5050", "Port to serve webserver on")
	browser := flag.Bool("browser", true, "open up the html interface in the default web browser")
	pluginList := os.Getenv("SEEKR_PLUGINS")
	plugins := []string{}
	if pluginList != "" {
		plugins = strings.Split(pluginList, ",")
	}
	flag.Parse()

	apiConfig, err := api.ApiConfig{
		Server: server.Server{
			Ip:        *ip,
			Port:      uint16(*port),
			ForcePort: *forcePort,
			WebServer: webserver.Webserver{
				Disable:    !*enableWebserver,
				FileSystem: content,
			},
			ApiServer: server.ApiServer{
				Disable: !enableApiServer,
			},
		},
		LogFile:       "seekr.log",
		DataBaseFile:  *data,
		DataBase:      dataBase,
		SetCORSHeader: true,
		SaveDBFunc:    api.DefaultSaveDB,
		LoadDBFunc:    api.DefaultLoadDB,
	}.ConfigParse()
	if err != nil {
		log.Panicf("error: %s", err)
	}
	apiConfig, err = seekrplugin.Open(plugins, apiConfig)
	if err != nil {
		log.Panicf("error: %s", err)
	}

	if *browser && !apiConfig.Server.WebServer.Disable {
		openbrowser(fmt.Sprintf("http://%s:%d/web/index.html", apiConfig.Server.Ip, apiConfig.Server.Port))
	}
	//fmt.Println("Welcome to seekr a powerful OSINT tool able to scan the web for " + strconv.Itoa(len(api.DefaultServices)) + "services")
	go api.Seekrd(api.DefaultSeekrdServices, 30) // run every 30 minutes
	api.ServeApi(apiConfig)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	api.Check(err)
}
