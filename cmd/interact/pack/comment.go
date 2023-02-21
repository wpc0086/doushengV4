package pack

import (
	"doushengV4/cmd/interact/dal/db"
	"doushengV4/kitex_gen/interact"
)

// Comment pack Comment info
func Comment(c *db.Comment) *interact.Comment {
	if c == nil {
		return nil
	}

	return &interact.Comment{Id: c.Id, Content: c.Content, CreateDate: c.CreateDate}
}

// Comments pack list of Comment info
func Comments(us []*db.Comment) []*interact.Comment {
	Comments := make([]*interact.Comment, 0)
	for _, c := range us {
		if temp := Comment(c); temp != nil {
			Comments = append(Comments, temp)
		}
	}
	return Comments
}
