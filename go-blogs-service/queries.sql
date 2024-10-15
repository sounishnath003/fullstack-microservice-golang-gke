-- queries.sql

-- name: getLatestRecommendedBlogs
-- some: param
-- other: param
SELECT ID, UserID, Title,SubTitle,
Content,CreatedAt,UpdatedAt 
FROM blogs ORDER BY CreatedAt DESC LIMIT 9;

-- name: createNewBlogpost
-- some: param
-- other: param
INSERT INTO blogs (UserID, Title, SubTitle, Content)
VALUES ($1,$2,$3,$4);

-- name: getBlogsByUsername
SELECT ID, UserID, Title, SubTitle, Content, CreatedAt, UpdatedAt
FROM blogs WHERE UserID=$1;

-- name: getBlogsByBlogID
SELECT ID, UserID, Title, SubTitle, Content, CreatedAt, UpdatedAt
FROM blogs WHERE ID=$1;
