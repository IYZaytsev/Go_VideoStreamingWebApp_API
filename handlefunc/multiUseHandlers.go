package handlefunc

import "net/http"

func Test(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "randomVideo.mp4")
}
