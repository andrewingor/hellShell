//
package main

import (
	"net/http"
	"io"
	//"fmt"
	//"os/exec"

	//"github.com/andrewingor/hellShell/go"
)
//Revision
const revision = "$Id$"
//
func init () {
//
}
// conclusion of contract
func conclusion (
		w http.ResponseWriter,
		req *http.Request	) {

	io.WriteString( w, webmuzzle )
}
// Hell server Run and Contracts conclusion
func main() {
	workdir := "."

	http.HandleFunc ("/", conclusion )
	//http.HandleFun("/put", filePut)
	//http.HandleFun("/get", fileGet)

	http.ListenAndServe(":1666", nil )
	http.FileServer(http.Dir (workdir))
}

var webmuzzle string = `
<!DOCTYPE html>
<html>
<head><title>Hell$hell</title></head>
<style>
body {
	text-align: center;
	font-family: Consolas;
	font-size: 20pt
}
.center {
	height: 7em;
	line-height: height
}
</style>
<body>
<form class="center">
<input type="text" name="args" value="cmd.exe wait arguments here" />
<input type="submit" value="Run" /><br/>
</form>
</body>
</html>
`
//EOF