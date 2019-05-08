func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method!="POST"{
		http.ServeFile(w,r,"signup.html")
		return
	}
	username:=r.FormValue("username")
	password:=r.FormValue("password")

	var user string

	err:=db.QueryRow("SELECT username FROM User_Login_Database WHERE username=?",username).Scan(&user)

	switch  {
	case err==sql.ErrNoRows:
		hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err!=nil{
			http.Error(w,"Unable to crate accout right now",http.StatusInternalServerError)
			return
		}
		_,err=db.Exec("INSERT INTO users(username,password) VALUES(?,?)",username,hashedPassword)
		if err!=nil{
			http.Error(w,"Unable to create your account",http.StatusInternalServerError)
			return
		}
		r.Write([]byte("User Account Created!"))


	case err!=nil:
		http.Error(w,"Unable to create your account",http.StatusInternalServerError)
		return
	default:
		http.Redirect(w,r,"/",http.StatusTemporaryRedirect)

	}

}
