package main

import (
	"net/http"

	"github.com/DAlba-sudo/pff"
)

func main() {
	// This sets up the web server. You generally do not need to modify this block.
	// It tells the server to look for your HTML files in the "./templates/" folder.
	app := pff.CreateApp(pff.Configuration{
		TemplateDirectoryPath: "./templates/",
		Address:               "0.0.0.0",
		Port:                  8080,
	})

	// FRONTEND DEVELOPERS: To add a new web page, copy and modify the block below.
	// The first parameter ("/some-random-endpoint") is the URL the user visits in their browser.
	// The second parameter ("./templates/random") is the path to your specific HTML file.
	random := app.RegisterTemplate("/some-random-endpoint", "./templates/random", pff.TemplateRegistrationOpts{
		// Set to true if this page should be wrapped inside your shared base layout (e.g., base.html).
		IncludeBaseTemplate: true,

		// Set to true if anyone can view this page without needing to log in.
		IsSessionUnprotected: true,
	})

	// FRONTEND DEVELOPERS: This connects backend data to your HTML file.
	// The string "Profile" is the variable name you will use in your HTML to display data.
	// For example, you can write {{ .Profile.Name }} in your HTML to show the name.
	// The backend engineer will build and provide the "BridgeProfile{}" component.
	random.RegisterBridge("Profile", BridgeProfile{})

	// This boots up the server. You do not need to touch this.
	err := app.Start()
	if err != nil {
		panic(err)
	}
}

// =========================================================================
// BACKEND DEVELOPERS: Everything below this line is your responsibility.
// Frontend developers do not need to edit or worry about the code below.
// =========================================================================

// Profile defines the structure of data sent to the frontend HTML template.
type Profile struct {
	Name  string
	Email string
}

// BridgeProfile contains the logic to fetch data from a database or API.
type BridgeProfile struct{}

func (b BridgeProfile) Data(wr http.ResponseWriter, req *http.Request) (any, error) {
	// The backend engineer writes the logic here to securely fetch the data
	// and return it so the frontend can display it.
	profile := &Profile{
		Name:  "George P. Burdell",
		Email: "George.Burdell@gatech.edu",
	}

	return profile, nil
}

