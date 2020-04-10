package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"./canvas"
	"./canvas/model"
)

func main() {
	subdomain := ""
	if len(os.Args) >= 2 {
		subdomain = os.Args[1]
	}
	token := ""
	if len(os.Args) == 2 {
		var err error
		var b []byte
		if b, err = ioutil.ReadFile(fmt.Sprintf("%s.pri", subdomain)); err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Error: no authentication token specified but file '%s.pri' does not exist.\n", subdomain)
			}
		}
		token = string(b)
	} else {
		token = os.Args[2]
	}
	if len(subdomain) == 0 || len(token) == 0 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <canvas subdomain> [authenication token]\n"+
			"\n"+
			"canvas subdomain:    This is the subdomain of instructure.com to use.\n"+
			"                     For example, this would be 'canvas' for the domain 'canvas.instructure.com'.\n"+
			"\n"+
			"authenication token: This is the authentication token to use for connecting to Canvas.\n"+
			"                     If this argument is not given, a file named <canvas subdomain>.pri must be\n"+
			"                     present in the current directory that contains the token.\n", os.Args[0])
		os.Exit(1)
	}
	c, err := canvas.CreateCanvas(strings.TrimSpace(subdomain), strings.TrimSpace(token))
	if err != nil {
		panic(err)
	}
	if err = c.Request("courses", nil, &[]model.Course{}, func(obj interface{}) error {
		courses := obj.(*[]model.Course)
		fmt.Println(courses)
		return nil
	}); err != nil {
		panic(err)
	}
}
