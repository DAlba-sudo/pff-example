package main

import (
	"net/http"

	"github.com/DAlba-sudo/pff"
)

func main() {
	app := pff.CreateApp(pff.Configuration{
		TemplateDirectoryPath: "./templates/",
		Address:               "0.0.0.0",
		Port:                  8080,
	})

	random := app.RegisterTemplate("/some-random-endpoint", "./templates/random", pff.TemplateRegistrationOpts{
		// use this option when you're building a page with other
		// components
		IncludeBaseTemplate: true,

		// use this option if the user does not need to be signed in to
		// use this component
		IsSessionUnprotected: true,
	})

	// This statement says that prior to rendering the template,
	// we will go and get a user Profile.
	random.RegisterBridge("Profile", BridgeProfile{})

	err := app.Start()
	if err != nil {
		panic(err)
	}
}

// A "Bridge" is a function that interacts with the backend and returns
// some data struture with content from that API (i.e., a Profile with name
// and email, etc).

type Profile struct {
	// this is an example of a "class" that
	// we can use to hold data in.
	Name  string
	Email string
}

type BridgeProfile struct{}

func (b BridgeProfile) Data(wr http.ResponseWriter, req *http.Request) (any, error) {
	// This is an example of a "bridge" that returns a
	// profile object.

	// Example of Request Approach to Populating Struct.
	//
	// backend_api_request, err := http.NewRequest("GET", "https://some-backend-api.com/profile", nil)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// resp, err := http.DefaultClient.Do(backend_api_request)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// var profile *Profile
	// if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
	// 	return nil, err
	// }

	// Here we are manually setting the data. In reality,
	// we would do something like what is commented above...
	profile := &Profile{
		Name:  "George P. Burdell",
		Email: "George.Burdell@gatech.edu",
	}

	return profile, nil
}
