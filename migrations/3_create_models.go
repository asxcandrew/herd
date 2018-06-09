package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating tables...")
		_, err := db.Exec(`CREATE TABLE tags (
												id serial PRIMARY KEY,
												name varchar UNIQUE,
												count int DEFAULT 0
											);
											CREATE TABLE topics (
												id serial PRIMARY KEY,
												name varchar,
												uri varchar UNIQUE,
												caption varchar,
												active boolean
											);
											CREATE TABLE related_topics (
												id serial PRIMARY KEY,
												related_id int REFERENCES topics(id),
												topic_id int REFERENCES topics(id)
											);
											CREATE TABLE stories (
												id serial PRIMARY KEY,
												uid char(12) UNIQUE,
												title varchar,
												subtitle varchar,
												description varchar,
												author_id int REFERENCES users(id),
												active boolean,
												published_at timestamptz,
												created_at timestamptz DEFAULT current_timestamp
											);
											CREATE TABLE story_tags (
												id serial PRIMARY KEY,
												story_id int REFERENCES stories(id),
												tag_id int REFERENCES tags(id)
											);
											CREATE TABLE story_revision (
												id serial PRIMARY KEY,
												story_id int REFERENCES stories(id),
												markdown varchar
											);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping tables...")
		_, err := db.Exec(`
				DROP TABLE tags;
				DROP TABLE topics;
				DROP TABLE stories;
				DROP TABLE story_tags;
				DROP TABLE story_revision;
				DROP TABLE related_topics;
			`)
		return err
	})
}
