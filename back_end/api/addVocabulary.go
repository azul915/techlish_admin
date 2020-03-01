package vocabulary

import (
	"time"
)

type Vocabulary struct {
	Word      string
	Category  string
	Mean      string
	Any       string
	CreatedAt time.Time
}

func AddVocabulary() string {
	return "ok"
}
