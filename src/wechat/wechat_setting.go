package wechat

import "fmt"

const APP_ID = "wxc6ae02f7514a9a33"
const APP_SECRET = "0117c4dfa79fb058177deea13e8c025d"
const TOKEN = "token123"

var TOKEN_URL = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",APP_ID, APP_SECRET)

const STAFF_ADD_URL = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"
const STAFF_UPDATE_URL = "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s"
const STAFF_DELETE_URL = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s"
const STAFF_UPLOAD_HEAD_IMAGE_URL = "http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%S&kf_account=%S"
const STAFF_LIST_URL = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s"
const STAFF_SEND_URL = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"

