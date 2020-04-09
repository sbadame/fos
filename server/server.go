package main

import (
	"io"
	"net/http"
	"os/exec"
	"flag"
	"fmt"
	"log"
	"strconv"
)

var (
	port = flag.Int("port", 8080, "Port to run on.")
	bin = flag.String("bin", "/bin/echo", "Binary to run on toggle.")
)

func main() {
	flag.Parse()

	state := "ON"
    http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("/bin/echo", fmt.Sprintf("GPIO. 1 %s", state))
		if stdoutStderr, err := cmd.CombinedOutput(); err != nil {
			io.WriteString(w, fmt.Sprintf("Error from %s: %v", *bin, err))
		} else {
			w.Write(stdoutStderr)
		}

		if state == "ON" {
			state = "OFF"
		} else {
			state = "ON"
		}
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, `
            <!DOCTYPE html>
            <head>
  				<title>Light</title>
				  <!-- This is a responsive site, don't make browsers lie about their size... -->
				  <meta charset="utf-8" name="viewport" content="width=device-width,initial-scale=1">
				  <style>
				  </style>
				  <script>
					function toggle() {
						fetch("/toggle");
					}
				  </script>
			</head>
			<body>
				<button onclick="toggle()">Toggle</button>
			</body>
        `)
    })
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(*port), nil))
}