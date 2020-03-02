package vocabulary

import (
	"fmt"
	"time"
)

type Vocabulary struct {
	Word      string
	Category  string
	Mean      string
	Any       string
	CreatedAt time.Time
}

func AddVocabulary(v *Vocabulary) {
	fmt.Println(v.Word)
	fmt.Println(v.Category)
	fmt.Println(v.Mean)

}
