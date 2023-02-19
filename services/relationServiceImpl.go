package services

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/mapper"
	"github.com/RaymondCode/simple-demo/utils"
	"strings"
)

type RelationImpl struct {
}

func (u RelationImpl) Follow(userId string, toUserId string, actionType string) {
	// 更新关心用户列表
	userFollowFilePath := config.UserFollowPrefixPath + userId
	followers := utils.ReadFollowers(userFollowFilePath)
	fmt.Print(followers + "\n")
	fmt.Print(actionType + "\n")
	if actionType == "1" {
		followers += fmt.Sprintf("-%s-", toUserId)
	} else {
		followers = strings.Trim(followers, fmt.Sprintf("-%s-", toUserId))
	}
	fmt.Print(followers + "\n")
	utils.WriteResult(followers, userFollowFilePath)

	// 更新粉丝列表
	fans := utils.ReadFollowers(config.UserFollowPrefixPath + toUserId + "-followed")
	if actionType == "1" {
		fans += fmt.Sprintf("-%s-", userId)
	} else {
		fans = strings.Trim(fans, fmt.Sprintf("-%s-", userId))
	}
	fmt.Print(fans + "\n")
	utils.WriteResult(fans, config.UserFollowPrefixPath+toUserId+"-followed")
}

func (u RelationImpl) FollowList(userId string) []entity.UserInfo {
	var userFollowList []entity.UserInfo
	follows := utils.ReadFollowers(config.UserFollowPrefixPath + userId)
	for i := 1; i < len(follows); i += 3 {
		j := i
		for {
			if j >= len(follows) || follows[j] != '-' {
				break
			}
			j++
		}
		followI := follows[i:j]
		fmt.Printf("粉丝id为%s", followI)
		userInfo := mapper.SelectUserInfoById(followI)
		userFollowList = append(userFollowList, *userInfo)
	}
	return userFollowList
}

func (u RelationImpl) FollowerList(userId string) []entity.UserInfo {
	var userFollowList []entity.UserInfo
	follows := utils.ReadFollowers(config.UserFollowPrefixPath + userId + "-followed")
	for i := 1; i < len(follows); i += 3 {
		j := i
		for {
			if j >= len(follows) || follows[j] == '-' {
				break
			}
			j++
		}
		followI := follows[i:j]
		fmt.Printf("粉丝id为%s", followI)
		userInfo := mapper.SelectUserInfoById(followI)
		userFollowList = append(userFollowList, *userInfo)
	}
	return userFollowList
}

func (u RelationImpl) FriendList(userId string) []entity.UserInfo {
	return u.FollowerList(userId)
}
