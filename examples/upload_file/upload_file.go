package examples

import (
	"github.com/kilfu0701/gogae/upload"
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
Generate Upload URL API => <a href="{{.Data["uploadEntry"]}}">{{.Data["uploadEntry"]}}</a>
</body></html>
`
)

func init() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc(uploadEntry, handleGenerateUrl)

	settings := upload.Settings{
		Bucket:  "asd",
		Folder:  "customer/photo",
		MaxSize: 1024 * 1024 * 10,    // 10MB
		BlobUrl: uploadEntry,
	}

	url, _ := upload.GenerateUploadURL(ctx, settings)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	w.Header().Set("Content-Type", "text/html")

	var indexTemplate = template.Must(template.New("index").Parse(indexTemplateHTML))
	var Data map[string]interface{}

	Data["uploadEntry"] = uploadEntry

	err = indexTemplate.Execute(w, Data)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
}

func handleGenerateUrl(w http.ResponseWriter, r *http.Request) {

}
