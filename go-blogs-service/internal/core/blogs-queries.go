package core

import "database/sql"

type BlogsServiceQueries struct {
	GetLatestRecommendedBlogs *sql.Stmt `query:"getLatestRecommendedBlogs"`
	CreateNewBlogpost         *sql.Stmt `query:"createNewBlogpost"`
	GetBlogsByUserID          *sql.Stmt `query:"getBlogsByUserID"`
	GetBlogsByBlogID          *sql.Stmt `query:"getBlogsByBlogID"`
}
