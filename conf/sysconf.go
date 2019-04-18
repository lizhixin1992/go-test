package conf

import "time"

const ViewDir = "./web/views"
const ViewExtension = ".html"
const HtmlReload = true

const SysTimeForm = "2006-01-02 15:04:05"
const SysTimeFormShort = "2006-01-02"

const SessionExpires = 30 * time.Minute

const StaticAssets = "./web/public/"
const Favicon = "favicon.ico"

var SysTimeLocation, _ = time.LoadLocation("Asia/shanghai")