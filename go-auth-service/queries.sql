-- queries.sql

-- name: getUserByID
-- some: param
-- some_other: param
SELECT ID, FirstName, LastName, Username, Email, Password
FROM users WHERE ID=$1;

-- name: createNewUser
-- some: param
-- some_other: param
INSERT INTO users (FirstName, LastName, Username, Email, Password) 
VALUES ($1,$2,$3,$4,$5);

-- name: addUserToUserRole
INSERT INTO user_roles (UserID, RoleID) VALUES ($1,$2);

-- name: getUserByUsername
-- some: param
-- some_other: param
SELECT ID, FirstName, LastName, Username, Email, Password
FROM users WHERE Username=$1;

-- name: getUserForJWT
-- some: param
-- some_other: param
SELECT u.FirstName, u.LastName, u.Username, u.Email, u.Password, r.Role
FROM users AS u JOIN user_roles AS ur ON u.id = ur.userid
JOIN roles AS r ON ur.roleid=r.id
WHERE u.Username=$1;
