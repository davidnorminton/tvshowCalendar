package web

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func StartWebServer() {
	port := getPort()
	openBrowser("http://127.0.0.1:" + port)
	path, _ := os.Getwd()

	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/search", SearchHandler)
	mux.HandleFunc("/details", DetailsHandler)
	mux.HandleFunc("/addshow", AddShowHandler)
	mux.HandleFunc("/removeshow", RemoveShowHandler)
	mux.HandleFunc("/updatefile", UpdateHandler)
	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir(path + "/web/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}

//getPort finds the next open port after 8080 to run the web interface
func getPort() string {
	port := 8080
	for {
		ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			port++
			continue
		}

		ln.Close()
		fmt.Println("Opening server on port: " + strconv.Itoa(port))
		return strconv.Itoa(port)
	}

}

// open opens the specified URL in the default browser of the user.
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
