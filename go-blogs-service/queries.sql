-- queries.sql

-- name: getLatestRecommendedBlogs
-- some: param
-- other: param
SELECT ID, UserID, Title,SubTitle,
Content,CreatedAt,UpdatedAt 
FROM blogs ORDER BY CreatedAt DESC LIMIT 10;

-- name: createNewBlogpost
-- some: param
-- other: param
INSERT INTO blogs (UserID, Title, SubTitle, Content)
VALUES ($1,$2,$3,$4);

