package config

// network
var ip = "192.168.110.107"
var port = "8080"
var addrPre = "http://" + ip + ":" + port + "/"

// video
var LocalVideoPrefixPath = "./public/videos/"
var LocalCoverPrefixPath = "./public/covers/"
var PlayUrlPrefix = addrPre + "public/videos/"
var CoverUrl = addrPre + "public/covers/"

// user
var UserFollowPrefixPath = "./userDatas/"
var UserAvatarPrefixPath = addrPre + "public/avatars/"
var UserBackgroundPrefixPath = addrPre + "public/background/"
