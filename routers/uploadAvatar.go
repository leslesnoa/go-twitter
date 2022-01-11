package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extention = strings.Split(handler.Filename, ".")[1]
	var record string = "uploads/avatars/" + IDUserInfo + "." + extention

	f, err := os.OpenFile(record, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error when uploading image! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error when copy image! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.UserInfo
	var status bool

	user.Avatar = IDUserInfo + "." + extention
	status, err = db.ModifyRecord(user, IDUserInfo)
	if err != nil || status == false {
		http.Error(w, "Error when saving the avatar in the DB! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
