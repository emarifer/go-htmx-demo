package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var tmpl *template.Template

/* var funcMap = template.FuncMap{
	"equal": func(n int) bool { return n == 5 },
	"inc":   func(n int) int { return n + 1 },
} */

/* templates will be parsed once at package first import */
func init() {
	if tmpl == nil {
		if tmpl == nil {
			tmpl = template.Must(tmpl.ParseGlob("views/layouts/*.html"))
			template.Must(tmpl.ParseGlob("views/*.html"))
			template.Must(tmpl.ParseGlob("views/partials/*.html"))
		}
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
	var intPage int
	intPage, _ = strconv.Atoi(r.URL.Query().Get("page"))
	if intPage == 0 {
		intPage = 1
	}

	offset := (intPage - 1) * 5

	note := new(Note)
	notesSlice, err := note.GetAllNotes(offset)
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	convertedNotes := []ConvertedNote{}
	for _, note := range notesSlice {
		newConvertedNote := convertDateTime(note, r.Header.Get("X-TimeZone"))
		convertedNotes = append(convertedNotes, newConvertedNote)
	}

	data := map[string]any{
		"Notes":    convertedNotes,
		"IncPage":  intPage + 1,
		"ShowMore": len(convertedNotes) == 5,
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
	_, err := newNote.CreateNote()
	if err != nil {
		var message string

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			message = "The title is already in use 🔥!"
		} else if strings.Contains(err.Error(), "CHECK constraint failed") {
			message = "The title is longer than 64 characters 🔥!"
		} else {
			message = fmt.Sprintf("Something went wrong: %s 🔥!", err)
		}

		w.Header().Set("HX-Retarget", "body")
		w.Header().Set("HX-Reswap", "beforeend")
		tmpl.ExecuteTemplate(w, "modal", message)

		return
	}

	w.Header().Set("HX-Location", "/")
}

func CompleteNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	note := new(Note)
	note.ID = id
	recoveredNote, err := note.GetNoteById()
	if err != nil {
		// w.Header().Set("HX-Trigger", "{\"myEvent\":\"The requested note was not found &#x1f631;!\"}")
		w.Header().Set("HX-Retarget", "body")
		w.Header().Set("HX-Reswap", "beforeend")
		tmpl.ExecuteTemplate(w, "modal", "The requested note was not found 😱!")

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
		w.Header().Set("HX-Retarget", "body")
		w.Header().Set("HX-Reswap", "beforeend")
		tmpl.ExecuteTemplate(w, "modal", "The requested note was not found 😱!")

		return
	}

	w.Header().Set("HX-Location", "/")
}

/* HOW TO EXTRACT URL QUERY PARAMETERS IN GO. VER:
https://freshman.tech/snippets/go/extract-url-query-params/

Parsear parámetros. VER:
https://www.sitepoint.com/get-url-parameters-with-go/
https://www.golangprograms.com/how-do-you-set-headers-in-an-http-response-in-go.html

ALTERNATIVE FORM FOR MODAL:
{{ define "modal" }}
<div id="modal"
    _="on closeModal add .closing then wait for animationend then remove me then reload() the location of the window end on myEvent from body put event.detail.value into #message then show me"
    style="display: none;">
    <div class="modal-underlay" _="on click trigger closeModal"></div>
    <div class="modal-content relative bg-base-100 p-6 rounded-2xl">
        <h3 class="font-bold text-lg">Go & HTMx Demo</h3>
        <p id="message" class="py-4"></p>

        <button _="on click trigger closeModal" class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
            ✕
        </button>
    </div>
</div>
{{ end }}
*/
