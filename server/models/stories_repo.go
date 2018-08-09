package models

//QueryStories returns query for index path
func QueryStories(stories *[]Story) error {
	return DB().Model(stories).Column("Author", "Tags").Select()
}
