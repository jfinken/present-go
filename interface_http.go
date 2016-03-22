/START2 OMIT
type MyEntity interface {
    UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, v MyEntity) error {
    return v.UnmarshalHTTP(r)
}

func (u *User) UnmarshalHTTP(r *http.Request) error {
    // Implement on our User object this specific method that allows the 
    // User to describe how it would get itself out of an HTTP request
    // ...
}

var u User
if err := GetEntity(req, &u); err != nil {
    // ...

    // This is safe because `var u User` will automatically
    // zero the User struct
}
/END2 OMIT

/START1 OMIT
func PostEntityHandler(w http.ResponseWriter, r *http.Request) {
    var b Banana
    if err := GetEntity(req, &b); err != nil {
        // ...
    }
    var ike Eisenhower 
    if err := GetEntity(req, &ike); err != nil {
        // ...
    }
    // ...
}

...

router.POST("/api/v1/entity", PostEntityHandler)
/END1 OMIT
