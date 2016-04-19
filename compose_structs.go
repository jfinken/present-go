/START1 OMIT
// Say we have a type User with some methods applicable to all Users:
type User struct {
  Name string
  Email string
}
func (u *User) SendEmail(subject, body string) bool {
  // ... send an email
  return true
}

// And an Admin, a User with additional capabilities
type Admin struct {
  *User
  SecretPassword string
}

func (a *Admin) UnlockGate() bool {
  // ... do admin stuff
  return true
}
/END1 OMIT

/START2 OMIT
// Call the User's Name property and methods explicitly
adm.User.Name 
adm.User.SendEmail("hi", "there")

// Or take advantage of type promotion 
adm.Name
adm.SendEmail("hi", "there")

// Likewise, you also have Admin's own abilities not available to type User
adm.UnlockGate()
/END2 OMIT
