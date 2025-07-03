package auth

//! PACKAGE SCOPE ⚠️
//email -> private
//Email -> Exporting

func extractSession() string{
	return "Session: LoggedIn ✅"
}

func GetSession()string{
	return extractSession()
}