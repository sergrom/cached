package main

import pb "github.com/sergrom/cached/pb"

func (c *Cached) getUsers() (map[int32]pb.User, error)  {
	users := make(map[int32]pb.User)

	const sqltpl = "SELECT id, name FROM user"

	rows, err := c.db.Query(sqltpl)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		user := pb.User{}

		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			panic(err)
		}

		users[user.Id] = user
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return users, nil
}