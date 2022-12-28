package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"pkg.cld19.com/mosbot/internal/client"
	"pkg.cld19.com/mosbot/core"
)

var (
	debug      bool
	docType    string
	docId      string
	configPath string
	copyToClip bool
	openURL    bool
)

func main() {
	debug = false

	docType = "doc"
	docId = "0"
	configPath = ""
	copyToClip = false
	openURL = false

	loadEnvVars()

	loadConfig()

	// command line args will override config settings
	argsv := os.Args[1:]

	if len(argsv) == 0 {
		displayHelp()
		os.Exit(0)
	}

	for i, v := range argsv {
		docid, _ := regexp.MatchString("^[0-9]+.[0-9]$", v)
		if docid {
			v = "id"
		}
		switch v {
		case "id":
			docId = argsv[i]
		case "doc":
		case "bug":
		case "patch":
		case "idea":
		case "sr":
			docType = v
		case "-o":
		case "--open":
			openURL = true
		case "-c":
		case "--copy":
			copyToClip = true
		case "-h":
		case "--help":
		case "help":
		default:
			displayHelp()
			os.Exit(0)
		}
	}

	if docId == "0" {
		displayHelp()
		os.Exit(0)
	}

	url, err := core.GenerateUrl(docType, docId)
	if err != nil {
		log.Fatal(err)
	}

	pubURL, _ := core.GeneratePublicUrl(docId)
	title, err := client.GetHtmlTitle(pubURL)

	if err != nil {
		log.Println(err)
		fmt.Println(url)
	} else {
		fmt.Printf("%s\n%s\n", title, url)
	}

	err = clipboard.WriteAll(url)
	if err != nil {
		log.Fatal(err)
	}
}

func loadEnvVars() {
	// should be full path to file
	// export MOS_CONFIG_PATH=/full/path/to/.mosbot.conf
	_, found := os.LookupEnv("MOS_CONFIG_PATH")
	if found {
		configPath = os.Getenv("MOS_CONFIG_PATH")
	}

	_, found = os.LookupEnv("MOS_COPY_URL")
	if found {
		copyToClip, _ = strconv.ParseBool(os.Getenv("MOS_COPY_URL"))
	}

	_, found = os.LookupEnv("MOS_OPEN_URL")
	if found {
		copyToClip, _ = strconv.ParseBool(os.Getenv("MOS_OPEN_URL"))
	}
}

func loadConfig() {
	configFile := ""
	// try to find and load config file
	if configPath != "" {
		configFile = configPath
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		configFile = fmt.Sprintf("%s/.mosbot.conf", homeDir)
	}

	// check if user can access file
	_, err := os.Open(configFile)
	if err != nil {
		if debug {
			if os.IsNotExist(err) {
				log.Println("Config file does not exist.")
			} else {
				log.Println("Cannot access config file (Maybe permissions?).")
			}
		}
	} else {
		parseConfigFile(configFile)
	}
}

func parseConfigFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var configKey string
	var configVal string
	var found bool

	scanner := bufio.NewScanner(file)

	// this will crash if line is over 64K
	for scanner.Scan() {
		configKey, configVal, found = strings.Cut(scanner.Text(), "=")
		if found {
			switch configKey {
			case "MOS_COPY_URL":
				copyToClip, _ = strconv.ParseBool(configVal)
			case "MOS_OPEN_URL":
				openURL, _ = strconv.ParseBool(configVal)
			}
			fmt.Println(configKey, configVal)
		} else {
			fmt.Println("Bad line in config file.")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func displayHelp() {
	msg := `mosbot - a utility to generate Oracle Support URLs

Usage:
mosbot <type> <id>

<type> is one of:
	* doc
	* bug
	* patch
	* idea
	* sr

	<id> is the document id, bug number, patch number or idea number

Options:
	-c    Copy the URL to the clipboard
	-o    Open the URL in a browser

Visit https://github.com/bmcculley/mosbot to learn how to use the configuration file.`

	fmt.Println(msg)
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
