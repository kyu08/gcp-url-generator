package config

import "fmt"

type bookmark struct {
	title string
	url   string
}

func (b bookmark) String() string {
	return fmt.Sprintf(
		`- title: '%s'
    url: '%s'
  `, b.title, b.url)
}
