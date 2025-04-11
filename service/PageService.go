package service

import (
	"database/sql"
	"math"
	"userManageSystem/models"
)

type Pagination struct {
	PageSize      int           `json:"page_size"`
	PerPage       int           `json:"per_page"`
	Total         int           `json:"total"`
	PageList      []Page        `json:"page_list"`
	UserList      []models.User `json:"user_list"`
	PrevPage      int           `json:"prev_page"`
	CurrentFilter int           `json:"current_filter"` // 当前过滤
	NextPage      int           `json:"next_page"`
	HasPrev       bool          `json:"has_prev"`
	HasNext       bool          `json:"has_next"`
}
type Page struct {
	Number int  `json:"number"`
	Active bool `json:"active"`
}

func List(limit, offset, status int, username string) Pagination {
	var resNum, resUser *sql.Rows
	// 避免走到查询禁用的用户
	if status == -1 && username == "" {
		resNum = models.GetUserNum()
		resUser = models.QueryPage(limit, offset)
	} else if username == "" {
		resNum = models.GetUserNumByStatus(status)
		resUser = models.QueryPageByStatus(limit, offset, status)
	} else {
		resNum = models.GetUserNumByName(username)
		resUser = models.QueryPageByName(limit, offset, username)
	}

	var num int
	for resNum.Next() {
		resNum.Scan(&num)
	}

	pageSize := offset
	perPage := int(math.Ceil(float64(num) / float64(pageSize)))

	var pageList []Page
	for i := 1; i <= perPage; i++ {
		if i == limit {
			pageList = append(pageList, Page{
				Number: i,
				Active: true,
			})
			continue
		}
		pageList = append(pageList, Page{
			Number: i,
			Active: false,
		})
	}

	var userList []models.User
	for resUser.Next() {
		var user models.User
		// 如果前面的字段查询为null，会导致后面的字段无法正确查询
		resUser.Scan(&user.Image, &user.Username, &user.Role, &user.Email, &user.Status, &user.LastLogin)
		user.Image = "/images/" + user.Image
		//user.Role = models.GetRoleNameByUserSession(user.SessionID)
		userList = append(userList, user)
	}

	return Pagination{
		PageSize:      pageSize,
		PerPage:       perPage,
		Total:         num,
		UserList:      userList,
		PageList:      pageList,
		PrevPage:      limit - 1,
		CurrentFilter: status,
		NextPage:      limit + 1,
		HasPrev:       limit > 1,
		HasNext:       limit < perPage,
	}
}
