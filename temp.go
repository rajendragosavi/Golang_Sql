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

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method!="POST"{
		http.ServeFile(w,r,"login.html")
		return
	}
	username:=r.FormValue("username")
	password:=r.FormValue("password")

	var dbusername string
	var dbpassword string

	err:= db.QueryRow("SELECT username, password FROM users WHERE username=?",username).Scan(&dbusername,&dbpassword)
	if err!=nil{
		http.Redirect(w,r,"/login",http.StatusTemporaryRedirect)
		return
	}
	err=bcrypt.CompareHashAndPassword([]byte(dbpassword),[]byte(password))
	if err!=nil{
		http.Redirect(w,r,"/login",http.StatusTemporaryRedirect)
		return
	}
	w.Write([]byte("Hello!"+dbusername))

}

func Homepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"index.html")
}

func Register(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"register.html")

	var VoterID int16
	var LastName string
	var FirstName string
	var State string
	var City string
	var Age string
	var Sex string

	_,err=db.Exec("INSERT INTO Voters(VoterID,LastName,FirstName,Age,Sex,State,City) VALUES(?,?,?,?,?,?,?,?)",VoterID,LastName,FirstName,Age,Sex,State,City)
	if err!=nil{
		http.Redirect(w,r,"/register",http.StatusTemporaryRedirect)
		return
	}
	w.Write([]byte(" Registered Successully!"))

}


