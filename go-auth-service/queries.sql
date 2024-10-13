-- queries.sql

-- name: getUserByID
-- some: param
-- some_other: param
SELECT Username, Password
FROM users WHERE ID=$1;

-- name: createNewUser
-- some: param
-- some_other: param
INSERT INTO users (FirstName, LastName, Username, Password) 
VALUES ($1,$2,$3,$4);

