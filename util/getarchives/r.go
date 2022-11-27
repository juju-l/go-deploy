package getarchives

//
///
func get() map[string]interface{} {
	var f = make(map[string]interface{})
	for _, o := range c {
		if o.DownloadUrl != "" {
			r, e := req().Get(o.DownloadUrl);if e!=nil { continue };f[o.Name] = /*string(*/r.Body()/*)*/
		}
	}
	return f
}

func init() {
	//
}