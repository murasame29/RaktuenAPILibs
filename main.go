package main

import (
	"fmt"
	"log"
	"server/rakutenapi"
)

func main() {
	body, err := rakutenapi.InitRequest().RakutenSearch.Get().Keyword("golang").Hits(1).GenreInformationFlag(false).Do()

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(body.Items[0])
	// hogehoge
}
