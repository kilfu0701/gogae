package examples

import (
	"net/http"
	"html/template"

	"github.com/kilfu0701/gogae/upload"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const (
	uploadEntry = "/api/photo/upload"
	uploadTemplateHTML = `
<html><body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br>
<input type="submit" name="submit" value="Submit">
</form></body></html>
`
	indexTemplateHTML = `
<html><body>
Generate Upload URL API => <a href="{{.UploadEntry}}">{{.UploadEntry}}</a>
</body></html>
`
	apiTemplateHTML = `
<html><body>
Your upload URL is => <b>{{.}}</b>
</body></html>
`
)

type Data struct {
	UploadEntry string
	UploadURL   string
}

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc(uploadEntry, handleGenerateUrl)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	w.Header().Set("Content-Type", "text/html")

	var indexTemplate = template.Must(template.New("index").Parse(indexTemplateHTML))

	data := Data{
		UploadEntry: uploadEntry,
	}

	err := indexTemplate.Execute(w, &data)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
}

func handleGenerateUrl(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	settings := upload.Settings{
		Bucket:  "asd",
		Folder:  "customer/photo",
		MaxSize: 1024 * 1024 * 10,    // 10MB
		BlobUrl: uploadEntry,
	}

	url, _ := upload.GenerateUploadURL(ctx, &settings)

	var apiTemplate = template.Must(template.New("api").Parse(apiTemplateHTML))
	err := apiTemplate.Execute(w, url)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
}
