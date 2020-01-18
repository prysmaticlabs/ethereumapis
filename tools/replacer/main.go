package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		panic("Usage: replacer template_file field1=file1 field2=other_file2")
	}

	t, err := template.ParseFiles(os.Args[1])
	if err != nil {
		panic(err)
	}

	index := make(map[string]template.HTML)
	for i := 2; i < len(os.Args); i++ {
		s := strings.Split(os.Args[i], "=")
		input, err := ioutil.ReadFile(s[1])
		if err != nil {
			panic(err)
		}
		index[s[0]] = template.HTML(template.JSEscapeString(string(input)))
	}

	if err := t.Execute(os.Stdout, index); err != nil {
		panic(err)
	}
}
