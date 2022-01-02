SELECT 
  u1.ID, u1.UserName, u2.UserName as ParentUserName 
FROM 
  USERS u1
  LEFT JOIN USERS u2 ON u1.Parent = u2.ID;