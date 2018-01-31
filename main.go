//The hell$hell is lightweight web file server for remote admins
//Browsing and transfering files, execute remote command
//By default serving at http://127.0.0.1:1666/

package main

import (
	//"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)
//TODO Git revision tag v2
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
//2TODO return $ favicon.ico recode from base64
			return
		}
		cmdstr := r.FormValue("cmd")
		filename := r.URL.Path
		if '/' != filename[len(filename)-1] && len(cmdstr) == 0 {
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
			if upfile, handler, err := r.FormFile("uploadfile"); err == nil {
				defer upfile.Close()
				if saveto, err := os.OpenFile(r.URL.Path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666); err == nil {
					defer saveto.Close()
					io.Copy(saveto, upfile)
				} //else
//TODO Error message panel to webmuzzle
			} 
		}

		htmlCmdForm := strings.Replace(htmlform, "$CMD$", cmdstr, 1)
		io.WriteString(w, htmlCmdForm) //Web muzzle

		if 0 < len(cmdstr) {
//TODO Set Timeout for exec or Kill button
//TODO Set Environment before execute 
/*
			cmd.Env = append(os.Environ(),
			    "FOO=duplicate_value",
			    "FOO=actual_value",
			)
*/
//1TODO Escaped: Space Into Name Trouble
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
	http.ListenAndServe("127.0.0.1:1666", nil)
}

//hell$hell run
func main() {}

//Web muzzle

//TODO HTML CSS to css
//TODO HTML CSS List of Style into webmuzzle & save to .ini
//TODO Auto setup meta codepage
var htmlhead string = `
<!DOCTYPE html>
<html>
<head><title>Hell$hell</title>
<meta codepage="$CODEPAGE$"/>
<meta =stylesheet content />
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
.promo {
	vertical-align: bottom;
	weight: 100%;
	font-family: Courier New;
	font-size: 9pt;
}
</style>
</head>
<body>
	<form enctype="multipart/form-data" action="" method="post">
 		<input type="file" name="uploadfile" />
 	   <input type="hidden" name="token" value="{{.}}"/>
 	  	 <input type="submit" value="upload" />
	</form>
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

//TODO Max Lince(R) site goo.glink
//TODO License goo.glink
//TODO table weight=100%
var htmltail string = `
<table class="promo">
<tr weight="100%">
<th><a target=_blank href="https://goo.gl/gVxGpd">License</a></th>
<th><a target=_blank href="https://github.com/andreingor/hellShell/">Source</a></th>
<td>Revision $Id$</td>
<th>&copy;2017-2018&nbsp;<a target=_blank href="https://goo.gl/CqgrAF">Max&nbsp;Lance(R)</a></th>
</tr>
</table>
</body>
</html>
`

//EOF
