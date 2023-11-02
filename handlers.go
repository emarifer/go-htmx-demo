package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var files = []string{
	"views/layouts/base.layout.html",
	"views/index.html",
	"views/about.html",
	"views/partials/button-up.html",
	"views/partials/footer.html",
	"views/partials/header.html",
	"views/partials/modal.html",
}

var tmpl *template.Template

/* templates will be parsed once at package first import */
func init() {
	if tmpl == nil {
		tmpl = template.Must(template.ParseFiles(files...))
	}
}

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	year := time.Now().Year()

	data := map[string]any{
		"Title": "Go & HTMx Demo",
		"Year":  year,
	}

	tmpl.ExecuteTemplate(w, "index.html", data)
}

func ShowAboutPage(w http.ResponseWriter, r *http.Request) {
	year := time.Now().Year()

	data := map[string]any{
		"Title": "About Me | Go & HTMx Demo",
		"Year":  year,
	}

	tmpl.ExecuteTemplate(w, "about.html", data)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(500 * time.Millisecond) // only to check how the spinner works

	// fmt.Println("Time Zone: ", r.Header.Get("X-TimeZone"))
	note := new(Note)
	notesSlice, err := note.GetAllNotes()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	convertedNotes := []ConvertedNote{}
	for _, note := range notesSlice {
		newConvertedNote := convertDateTime(note, r.Header.Get("X-TimeZone"))
		convertedNotes = append(convertedNotes, newConvertedNote)
	}

	data := map[string][]ConvertedNote{
		"Notes": convertedNotes,
	}

	tmpl.ExecuteTemplate(w, "note-list", data)
}

func AddNote(w http.ResponseWriter, r *http.Request) {

	title := strings.Trim(r.PostFormValue("title"), " ")
	description := strings.Trim(r.PostFormValue("description"), " ")
	if len(title) == 0 || len(description) == 0 {
		var errTitle, errDescription string
		if len(title) == 0 {
			errTitle = "Please enter a title in this field"
		}
		if len(description) == 0 {
			errDescription = "Please enter a description in this field"
		}

		data := map[string]string{
			"FormTitle":       title,
			"FormDescription": description,
			"ErrTitle":        errTitle,
			"ErrDescription":  errDescription,
		}

		w.Header().Set("HX-Retarget", "form")
		w.Header().Set("HX-Reswap", "innerHTML")
		tmpl.ExecuteTemplate(w, "new-note-form", data)

		return
	}

	newNote := new(Note)
	newNote.Title = title
	newNote.Description = description
	note, err := newNote.CreateNote()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	tmpl.ExecuteTemplate(w, "note-list-element", convertDateTime(note, r.Header.Get("X-TimeZone")))
}

func CompleteNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	note := new(Note)
	note.ID = id
	recoveredNote, err := note.GetNoteById()
	if err != nil {
		w.Header().Set("HX-Trigger", "{\"myEvent\":\"The requested note was not found &#x1f631;!\"}")
		return
	}

	updatedNote, err := recoveredNote.UpdateNote()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	tmpl.ExecuteTemplate(w, "note-list-element", convertDateTime(updatedNote, r.Header.Get("X-TimeZone")))
}

func RemoveNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	note := new(Note)
	note.ID = id
	err := note.DeleteNote()
	if err != nil {
		w.Header().Set("HX-Trigger", "{\"myEvent\":\"The requested note was not found &#x1f631;!\"}")
		return
	}
}

/* HOW TO EXTRACT URL QUERY PARAMETERS IN GO. VER:
https://freshman.tech/snippets/go/extract-url-query-params/

Parsear parámetros. VER:
https://www.sitepoint.com/get-url-parameters-with-go/
https://www.golangprograms.com/how-do-you-set-headers-in-an-http-response-in-go.html
*/
