SELECT a."ID", a."UserName", b."UserName" as "ParentUserName" 
FROM "user" as a
LEFT JOIN "user" as b ON b."ID" = a."Parent"