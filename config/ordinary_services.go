package config

import "fmt"

var servicesToBookmarksLength = len(services) * len(projects)

func servicesToBookmarks(projects []project, services []service) []bookmark {
	bookmarks := make([]bookmark, 0, servicesToBookmarksLength)
	for _, service := range services {
		for _, p := range projects {
			bookmarks = append(bookmarks, bookmark{
				title: fmt.Sprintf("%s %s", service.name, p.name),
				url:   fmt.Sprintf(service.url, p.key),
			})
		}
	}

	return bookmarks
}
