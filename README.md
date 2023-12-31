# Golang Full-stack Web App

#### Simple full-stack web application for saving notes to a Sqlite database (CRUD), with HTML template rendering using </> htmx & _hyperscript.

#### This is a minimalist application that does not use any additional libraries beyond the standard Go library, except for the necessary driver to manage the Sqlite database, thus following the trend of Golang developers to only use dependencies strictly necessary, taking advantage of all the power of the standard library.

#### Rendering is achieved by using the "html/template" package, i.e. Go's native form of rendering, and the "</> htmx" JavaScript library. The latter makes it possible to make requests to the backend (GET, POST, PATCH and DELETE) without reloading the page as in a SPA, but with a size of said library of only 15K. Additionally, "_hyperscript" is used, another JavaScript library developed by the same author with the purpose of performing a few actions by writing a kind of inline JavaScript code.

###### **Note**: the theoretical bases of Htmx and illustrative examples of its correct use can be consulted in the book HYPERMEDIA SYSTEMS written by the same developer of the library, Carson Gross, in [this site](https://hypermedia.systems/).

---

### Screenshots:

###### Homepage:

<img src="https://github.com/emarifer/go-htmx-demo/assets/68773736/16c11d53-3b0a-49a9-8f43-5fe22a2a1cf0" width="75%">

<br>

###### Error reporting modal:

<img src="https://github.com/emarifer/go-htmx-demo/assets/68773736/1b32d022-a765-40d5-aa83-65635d931134" width="75%">

---

### Setup:

Besides the obvious prerequisite of having Go! on your machine, you must have Air installed for hot reloading when editing code and NodeJs.

Since the application uses Tailwind as a CSS framework, you must run some NodeJs commands in the project root before starting the application:

```
$npm i
```

Next, whether you want to make code changes or create production CSS, you need to run these commands:

```
# If you want to edit the code and update the build CSS:

$npm run watch-css

# If you want to create the production CSS:

$npm run build-css-prod
```

Start the app in development mode:

```
$ air # Ctrl + C to stop the application
```

Build for production:

```
$ go build -ldflags="-s -w" -o ./bin/main . # ./bin/main to run the application
```

### Happy coding 😀!!
