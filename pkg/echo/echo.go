package echo

import (
	"fmt"

	corev1 "github.com/oif/go-skeleton/types/core/v1"
)

func Run() {
	fmt.Println(corev1.Echo{
		Name: "go",
	}.Hey())
}
