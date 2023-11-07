package main

import (
	"log"
	"net/http"
)

func init() {
	MakeMigrations()
}

func main() {

	http.HandleFunc("/", ShowHomePage)
	http.HandleFunc("/about", ShowAboutPage)
	http.HandleFunc("/notes", GetNotes)
	http.HandleFunc("/add-note", AddNote)
	http.HandleFunc("/update-note/", CompleteNote)
	http.HandleFunc("/delete-note/", RemoveNote)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("🚀 Starting up on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}

/* REFERENCES:
https://medium.com/@orlmonteverde/api-rest-con-go-golang-y-sqlite3-e378af30719c
https://medium.com/@back_to_basics/golang-template-1-bcb690165663
https://www.alexedwards.net/blog/working-with-cookies-in-go
https://lets-go.alexedwards.net/sample/02.07-html-templating-and-inheritance.html

https://htmx.org/attributes/hx-swap/
https://htmx.org/attributes/hx-target/
https://htmx.org/attributes/hx-boost/
https://hypermedia.systems/htmx-in-action/#_ajax_ifying_our_application
https://htmx.org/headers/hx-location/
https://htmx.org/essays/view-transitions/
https://htmx.org/docs/#special-events
https://htmx.org/attributes/hx-trigger/#standard-event-modifiers
https://htmx.org/examples/update-other-content/
https://www.jetbrains.com/guide/dotnet/tutorials/htmx-aspnetcore/out-of-band-swaps/
https://www.youtube.com/watch?v=g7Nlo6N2hAk
https://hyperscript.org/
https://hypermedia.systems/book/contents/

https://www.calhoun.io/intro-to-templates-p3-functions/
https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go
https://pkg.go.dev/text/template
https://golangforall.com/en/post/templates.html
https://krishbhanushali.medium.com/3-simple-techniques-using-go-templates-6718e2cc77e2
https://blog.logrocket.com/using-golang-templates/

https://github.com/moeenn/htmx-golang-demo
https://github.com/orlmonteverde/go-api-with-sqlite
https://github.com/NerdCademyDev/gophat
https://github.com/bugbytes-io/htmx-go-demo
https://github.com/awesome-club/go-htmx
https://github.com/marco-souza/marco.fly.dev
*/
