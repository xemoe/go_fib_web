package actions

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gobuffalo/httptest"
)

func (as *ActionSuite) Test_UploadHandler() {

	var tests = []struct {
		filename             string
		expectedStatusCode   int
		expectedUploadedFile string
		expectedResult       string
	}{
		{"test.jpg", http.StatusOK, "./uploads/test.jpg", "test.jpg"},
	}

	for _, t := range tests {

		os.RemoveAll("./uploads")

		r, err := os.Open(t.filename)
		as.NoError(err)

		form := struct {
			Name string
		}{"Foo"}
		wf := httptest.File{
			ParamName: "MyFile",
			FileName:  r.Name(),
			Reader:    r,
		}

		res, err := as.HTML("/api/v1/upload").MultiPartPost(form, wf)
		as.NoError(err)
		as.Equal(200, res.Code)

		//
		// Validate saved file
		//
		_, err = os.Stat(t.expectedUploadedFile)
		as.NoError(err)

		//
		// Validate json response message
		//
		var response Message
		json.Unmarshal(res.Body.Bytes(), &response)

		as.Equal(t.expectedStatusCode, res.Code)
		as.Equal(t.expectedResult, response.Result)

	}
}
