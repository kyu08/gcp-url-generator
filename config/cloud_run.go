package config

import "fmt"

var runServiceBookmarksLength = len(runServices) * len(projects)

func runServicesToBookmarks(runServices []string, region string, projects []project) []bookmark {
	bookmarks := make([]bookmark, 0, runServiceBookmarksLength)

	const urlFormat = "https://console.cloud.google.com/run/detail/%s/%s/revisions?project=%s"
	for _, runService := range runServices {
		for _, project := range projects {
			bookmarks = append(bookmarks, bookmark{
				title: fmt.Sprintf("Cloud Run %s %s", runService, project.name),
				url:   fmt.Sprintf(urlFormat, region, runService, project.key),
			},
			)
		}
	}

	return bookmarks
}
