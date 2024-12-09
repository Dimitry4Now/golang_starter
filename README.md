# Go Starter Project

This is a starter project for a website using **Golang**, **HTMX**, and **Tailwind CSS**. It's an extremely lightweight stack and is completely hypermedia-driven.

---

## About the Project

I've been planning on learning the Go programming language for a long time now, and I finally decided to do so. What better way to learn than by doing? I started this as a base for any future projects.

---

## File/Folder Structure

```go

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

```

---

### Folder Explanations

- **`db/`**  
  Contains all the database configurations and (eventually) data migrations. This is still a work in progress.

- **`handlers/`**  
  Contains all the individual routing handlers. Each handler is responsible for loading and serving the HTML templates. A `common.go` file is included to hold shared logic or components used across multiple handlers.

- **`i18n/`**  
  Holds the configuration files for localization (translation and multi-language support). Currently, it includes English and Macedonian (my native language) for demonstration purposes. Feel free to add more languages as needed.

- **`models/`**  
  This folder is for database entity models. Since this project is in its early stages, this part will evolve as I continue to learn Go.

- **`static/`**  
  Contains all static files used throughout the application, such as JavaScript and CSS. In the future, it will include images and other assets.

- **`templates/`**  
  Holds the HTML templates served to the client. These are styled using Tailwind CSS and leverage HTMX for sending requests to the Go backend.


- **`main.go`**  
  This is where everything comes together. The local http server is started and all the routes are handled, passing the functions from **`handlers/`**.
---

### Contribution

Feel free to use this project as a starting point for your own projects or contribute to improving this bare-bones Golang + HTMX web application starter.


