package config

import "fmt"

// TODO: implement gaeToString

var cloudSQLBookmarksLength = len(projects)

func cloudSqlBookmarks(projects []project) []bookmark {
	urlFormat := "https://console.cloud.google.com/sql/instances/%s/overview?project=%s"
	bookmarks := make([]bookmark, 0, cloudSQLBookmarksLength)
	for _, p := range projects {
		b := bookmark{
			title: fmt.Sprintf("Cloud SQL %s", p.name),
			url:   fmt.Sprintf(urlFormat, p.key, p.key),
		}
		bookmarks = append(bookmarks, b)
	}

	return bookmarks
}
