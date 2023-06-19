package version

var Notices = []string{
	"[Important] The `data` field in API Response will be deprecated and replaced with the `response` field. The data field is still present in this version of wdb-server for backward compatibility, and will be phased out in future version. Please ensure all your dependant systems move to use the `response` field instead of `data`. Refer this for more: https://github.com/TanmoySG/wunderDB/issues/121",
}
