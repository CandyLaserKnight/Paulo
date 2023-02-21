package paulo

import "net/http"

func (p *Paulo) SessionLoad(next http.Handler) http.Handler {
	p.InfoLog.Println("SessionLoad called")
	return p.Session.LoadAndSave(next)
}
