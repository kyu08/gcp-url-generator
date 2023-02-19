package config

func Main() []bookmark {
	bookmarks := make([]bookmark, 0, servicesToBookmarksLength*runServiceBookmarksLength*cloudSQLBookmarksLength)

	bookmarks = append(bookmarks, servicesToBookmarks(projects, services)...)
	bookmarks = append(bookmarks, runServicesToBookmarks(runServices, region, projects)...)
	bookmarks = append(bookmarks, cloudSqlBookmarks(projects)...)

	return bookmarks
}
