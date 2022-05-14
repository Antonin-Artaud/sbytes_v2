package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func main() {
	guid, _ := uuid.NewUUID()

	fmt.Println(guid)

	guidReplaced := strings.Replace(guid.String(), "-", "", -1)

	fmt.Println(guidReplaced)
}
