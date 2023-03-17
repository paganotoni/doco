package doco

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Build() error {
	fmt.Println("[INFO] Building docs")

	docs, err := Parse()
	if err != nil {
		return err
	}

	err = os.RemoveAll("public")
	if err != nil {
		return err
	}

	tmpl, err := templates.ReadFile("html/documents.html")
	if err != nil {
		return err
	}

	for _, doc := range docs {
		f := filepath.Dir(doc.ResultingPath())
		err = os.MkdirAll(f, 0777)
		if err != nil {
			return err
		}

		html, err := doc.HTML()
		if err != nil {
			return err
		}

		tt, err := template.New("file").Parse(string(tmpl))
		if err != nil {
			return err
		}

		bb := new(bytes.Buffer)
		tt.Execute(bb, struct {
			Content    template.HTML
			Navigation template.HTML
		}{
			Content:    template.HTML(html),
			Navigation: template.HTML(docs.NavigationHTML()),
		})

		fmt.Printf("[INFO] Write > %s\n", doc.ResultingPath())
		if err := ioutil.WriteFile(doc.ResultingPath(), bb.Bytes(), 0777); err != nil {
			return err
		}
	}

	return nil
}
