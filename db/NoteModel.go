package db

import "github.com/deta/deta-go/service/base"

type Note struct {
  Key string `json:"key"`
  UserKey string `json:"userKey"`

  Content string `json:"content"`
  Completed bool `json:"completed"`

  DateDay int `json:"dateDay"`
  DateMonth int `json:"dateMonth"`
  Time string `json:"time"`
}

func (note *Note) Put(userKey string) error {
  note.Key = GenKey()
  note.Completed = false
  note.UserKey = userKey

  _, err := Notes.Put(note)
  return err
}

func GetNotes(q base.Query) ([]Note, error) {
  var notes []Note

  _, err := Notes.Fetch(&base.FetchInput {
    Q: q, 
    Dest: &notes,
  })

  return notes, err
}

func ToggleComplete(key string) (Note, error) {
  var note Note
  Notes.Get(key, &note)

  var err error

  if (!note.Completed) {
    err = Notes.Update(key, base.Updates {
      "completed": true,
    })
  } else {
    err = Notes.Update(key, base.Updates {
      "completed": false,
    })
  }

  note.Completed = !note.Completed

  return note, err
}