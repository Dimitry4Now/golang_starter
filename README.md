## Go Starter Project

So this is a starter project for website using Golang combined with HTMX and Tailwind CSS. Extremely lightweight stack, and it is completely Hypermedia driven.

## About the project

I've been planing on learning the Go programming language for a long time now, and i finally decided to do so. And what better way of learning than learning by doing. So i started this as a base for any future projects including this stack.

The basic file/folder/project structure is as follows:

├── db
│   └── db.go
├── go.mod
├── go.sum
├── handlers
│   ├── aboutUs.go
│   ├── changeLanguage.go
│   ├── common.go
│   ├── contact.go
│   ├── landing.go
│   ├── login.go
│   └── reditectToLogin.go
├── i18n
│   ├── en.json
│   └── mk.json
├── main.go
├── models
│   └── user.go
├── static
│   ├── script.js
│   └── style.css
└── templates
    ├── aboutUs.html
    ├── contact.html
    ├── footer.html
    ├── header.html
    ├── landing.html
    └── login.html


where everything is separated in order for it to be easily managed. In the db folder will be all the database configuration and the data migrations(not yet added but i'll update it as i go). 

The handlers, as the name suggest, are the handlers for all the individual "routings". The logic that loads and serves the html templates. The commons.go file contains all the common components that are used in all of the handlers(and for the localization, all across the application).

In the folder named i18n are the configuration files for the localization(translation, multilanguage support) of the page. It has just English and Macedonian(my native) as options, just for demonstrating purposes. But feel free to add more as needed.

The models folder will contain all the models for the database entities, but this is a very early stage of "the project" and i am still new to Golang, so i will be updating here as well as i learn.

Static folder contains all the static files that will be needed throughout the whole application. For now contains only the .js and .css files, but it will eventually contain all the images and other stuff.

Under the templates folder, all the actual templates that are served to the client are stored. These are plain .html files, styled with tailwind CSS and use HTMX to send requests to Golang(the backend).

Feel free t0 use this project and feel free to contribute if anything is missing for a bare bone starters Golang+HTMX web application.
