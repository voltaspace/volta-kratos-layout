package rediskey

const (
	WxFosterTask              = "pan_wxfoster_task:"                 //.微信养号小程序个微任务 HASH
	WxworkFosterTask          = "pan_wxworkfoster_task:"             //.微信养号小程序企微任务 HASH
	WxFosterMobileSync        = "pan_wxfoster_mobile_sync"           //.微信养号小程序异步编辑账号 LIST
	WxFosterStationSync       = "pan_wxfoster_station_sync"          //.微信养号小程序异步编辑站点 LIST
	WxFosterTaskSync          = "pan_wxfoster_task_sync"             //.微信养号小程序异步编辑任务 LIST
	WxFosterMobileInfo        = "pan_wxfoster_mobile_info:"          //.微信养号小程序站点内账号列表 String
	WxFosterMobileBillInfo    = "pan_wxfoster_mobile_bill_info:"     //.微信养号小程序站点内账号账单列表 String
	WxFosterStationInfo       = "pan_wxfoster_station_info"          //.微信养号小程序站点信息 HASH
	WxFosterStationToken      = "pan_wxfoster_station_token:"        //.微信养号小程序站点登录token String
	WxFosterTeachingVideo     = "pan_wxfoster_teaching_video_info"   //.微信养号视频库
	WxFosterFriendInfo        = "pan_wxfoster_friend_info:"          //.微信养号小程序好友信息 HASH
	WxFosterMobileFriendId    = "pan_wxfoster_mobile_friend_id:"     //.微信养号小程序账号解封好友ID SET
	WxFosterCanSwitchFriendId = "pan_wxfoster_can_switch_friend_id:" //.微信养号站点内可被替换的解封微信好友ID SET
	WxFosterCanBindFriendId   = "pan_wxfoster_can_bind_friend_id:"   //.微信养号站点内可被绑定的解封微信好友ID SET
)
