package core

import "database/sql"

type BlogsServiceQueries struct {
	GetLatestRecommendedBlogs *sql.Stmt `query:"getLatestRecommendedBlogs"`
	CreateNewBlogpost         *sql.Stmt `query:"createNewBlogpost"`
	GetBlogsByUsername        *sql.Stmt `query:"getBlogsByUsername"`
}
