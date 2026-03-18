# Frontend Development Guide: Workflow and Templating

This document outlines the standard operating procedure for frontend engineers working with our custom Go web server configuration. The architecture strictly separates visual design from backend data retrieval, allowing you to build and style web pages using standard HTML and Bootstrap CSS.

## Part 1: The Development Workflow

**Step 1: Create the HTML Document**
Start by building your page using standard HTML and your preferred Bootstrap CSS components. Save this file inside the `templates` directory. You only need to focus on the specific content for this view, as the configuration handles wrapping your file in the global base layout automatically.

**Step 2: Register the Web Route**
Open the `main.go` file to make your page accessible on the web server. Add a new route using the `app.RegisterTemplate()` function. You will define the URL path the user types into their browser, provide the local path to your new HTML file, and set the boolean flags for base template inclusion and session protection.

**Step 3: Request Backend Data Bridges**
When your page requires dynamic content like user profiles or database records, collaborate with the backend team to build a Bridge. The backend engineer handles the security and data fetching, registering the final data structure to your route under a designated variable name, such as `Profile` or `Inventory`.

**Step 4: Render the Dynamic Content**
Return to your HTML file to connect the data. Using Go template syntax, you will inject the backend variables directly into your markup. This keeps your HTML clean while allowing the server to populate the exact details before sending the final page to the user.

## Part 2: Go Templating Syntax Reference

The server uses standard Go templating to pass data into your HTML. You can use the following syntax directly inside your tags and Bootstrap components to control how data is displayed.

| Concept | Syntax Example | Description |
| :--- | :--- | :--- |
| **Basic Variables** | `{{ .Profile.Name }}` | Outputs the exact text or number provided by the backend Bridge variable. |
| **If Statements** | `{{ if .Profile.Email }} <p>Email provided</p> {{ end }}` | Renders the enclosed HTML block only if the variable exists and is not empty. |
| **If/Else Statements** | `{{ if .IsAdmin }} <button>Edit</button> {{ else }} <span>View Only</span> {{ end }}` | Displays the first HTML block if the condition is true, and the second block if it is false. |
| **Loops** | `{{ range .Users }} <li>{{ .Name }}</li> {{ end }}` | Iterates over a list or array provided by the backend, repeating the enclosed HTML for every item in the list. |
| **Comments** | `{{/* This is a backend comment */}}` | Leaves a comment that is stripped out before the HTML is sent to the browser. |
