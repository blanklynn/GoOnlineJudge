package session

import (
    "net/http"

    "ojapi/model"
)

var SessionManager = NewManager()

func GetUser(r *http.Request) *model.User {
    sid := r.FormValue("access_token")

    if sid != "" {
        uid := GetSession(sid, "Uid")
        userModel := &model.UserModel{}
        user, err := userModel.Detail(uid)
        if err != nil {
            return nil
        }
        return user
    }
    return nil
}

func SetSession(sid string, key string, value string) {
    session := SessionManager.GetSession(sid)
    session.Set(key, value)
}

func GetSession(sid string, key string) string {
    sess := SessionManager.GetSession(sid)
    if sess == nil {
        return ""
    }

    return sess.Get(key)
}

func DeleteSession(r *http.Request) {
    sid := r.FormValue("access_token")
    SessionManager.DeleteSession(sid)
}

func StartSession() *Session {
    return SessionManager.StartSession()
}
