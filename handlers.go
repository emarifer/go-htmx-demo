package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var files = []string{
	"views/index.html",
	"views/partials/button-up.html",
	"views/partials/footer.html",
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-TimeZone")
	var timeZone string
	if err == nil {
		timeZone, _ = url.QueryUnescape(cookie.Value)
	}
	// fmt.Println("Time Zone: ", timeZone)

	note := new(Note)
	/* note.Title = "Primera Nota"
	note.Description = "Esta es la descripción de la Primera Nota"
	err := note.CreateNote()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	} */

	notesSlice, err := note.GetAllNotes()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	convertedNotes := []ConvertedNote{}
	for _, note := range notesSlice {
		newConvertedNote := convertDateTime(note, timeZone)
		convertedNotes = append(convertedNotes, newConvertedNote)
	}

	year := time.Now().Year()

	data := map[string]any{
		"Notes": convertedNotes,
		"Year":  year,
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, data)
}

func AddNote(w http.ResponseWriter, r *http.Request) {
	/* cookie, err := r.Cookie("X-TimeZone")
	var timeZone string
	if err == nil {
		timeZone, _ = url.QueryUnescape(cookie.Value)
	} */

	time.Sleep(1 * time.Second) // to check how the spinner works
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
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.ExecuteTemplate(w, "new-note-form", data)

		return
	}

	newNote := new(Note)
	newNote.Title = title
	newNote.Description = description
	_, err := newNote.CreateNote()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	// http.Redirect(w, r, "/", http.StatusOK)

	w.Header().Set("HX-Redirect", "/")

	/* tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl.ExecuteTemplate(w, "note-list-element", convertDateTime(note, timeZone)) */
}

func CompleteNote(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-TimeZone")
	var timeZone string
	if err == nil {
		timeZone, _ = url.QueryUnescape(cookie.Value)
	}

	urlStr := r.URL.String()
	myUrl, _ := url.Parse(urlStr)
	params, _ := url.ParseQuery(myUrl.RawQuery)
	id, _ := strconv.Atoi(params.Get("id"))

	// fmt.Println("ID: ", id)

	note := new(Note)
	note.ID = id
	recoveredNote, err := note.GetNoteById()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	updatedNote, err := recoveredNote.UpdateNote()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl.ExecuteTemplate(w, "note-list-element", convertDateTime(updatedNote, timeZone))
}

func RemoveNote(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.String()
	myUrl, _ := url.Parse(urlStr)
	params, _ := url.ParseQuery(myUrl.RawQuery)
	id, _ := strconv.Atoi(params.Get("id"))

	// fmt.Println("ID: ", id)
	note := new(Note)
	note.ID = id
	err := note.DeleteNote()
	if err != nil {
		log.Fatalf("something went wrong: %s", err.Error())
	}

	/* w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusNoContent) */
}

/* Parsear parámetros. VER:
https://www.sitepoint.com/get-url-parameters-with-go/
https://www.golangprograms.com/how-do-you-set-headers-in-an-http-response-in-go.html
*/
