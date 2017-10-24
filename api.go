package facebook

import "bytes"

type (
	// API contains information that can be passed
	// to access facebook graph api
	API struct {
		FB   *Facebook
		Path string

		SelectedFields string
		Message        string
		ShareURL       string
	}

	// APIInterface is the main functionality of facebook-go-sdk
	APIInterface interface {
		Fields(fields string) APIInterface
		Messages(msg string) APIInterface
		Link(shareURL string) APIInterface

		Get() (interface{}, error)
		Post() (interface{}, error)
		Delete() (interface{}, error)
	}
)

// Fields is use to set selected fields to be return
// by facebook graph api
func (a *API) Fields(fields string) APIInterface {
	a.SelectedFields = fields
	return a
}

// Messages is use to define message that want to be posted
// into user's timeline
func (a *API) Messages(msg string) APIInterface {
	a.Message = msg //`{"message": "` + msg + `"}`
	return a
}

// Link is use to define a share link that want to be posted
// into user's timeline
func (a *API) Link(shareURL string) APIInterface {
	a.ShareURL = shareURL
	return a
}

// Get is use to submit request using GET method
func (a *API) Get() (interface{}, error) {
	url := new(Graph).
		SetVersion(a.FB.GetVersion()).
		GenerateURLWithFields(a.Path, a.SelectedFields)

	return new(Call).
		SetMethod(`GET`).
		SetURL(url).
		AddHeader(`Authorization`, a.FB.AccessToken).
		Submit()
}

// Post is use to submit request using POST method
func (a *API) Post() (interface{}, error) {
	url := new(Graph).
		SetVersion(a.FB.GetVersion()).
		GenerateURL(a.Path)

	body := `{"message": "` + a.Message + `", "link": "` + a.ShareURL + `"}`

	return new(Call).
		SetMethod(`POST`).
		SetURL(url).
		SetBody(bytes.NewBuffer([]byte(body))).
		AddHeader(`Authorization`, a.FB.AccessToken).
		AddHeader(`Content-Type`, `application/json`).
		Submit()

}

// Delete is use to submit request using DELETE method
func (a *API) Delete() (interface{}, error) {
	url := new(Graph).
		SetVersion(a.FB.GetVersion()).
		GenerateURL(a.Path)

	return new(Call).
		SetMethod(`DELETE`).
		SetURL(url).
		AddHeader(`Authorization`, a.FB.AccessToken).
		Submit()
}
