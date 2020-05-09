# ext-errors
Package `exterrors` provides extended error handling primitives to add a bit more info to errors returning from the function.

## Motivation
In every API that handles HTTP requests, we need to work with error handling.
We'd like to add extra information about the response status and can't figure it out with the standard `errors` package.
With `exterrors` package, you know the status of your request in any place of the application.

## Usage
Use `ExtError` interface instead of standard `error`

```go
func GetUserByID(id string) (User, exterrors.ExtError) {
    user, err := db.GetUserByID(id)
    if err != nil {
        return User{}, exterrors.NewInternalServerError(err)
    }
    
    return user, nil
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")    
    
    user, extErr := GetUserByID("user_id_1")
    if extErr != nil {
        SendExtError(w, extErr)
        return
    }
}

func SendExtError(w http.ResponseWriter, extErr exterrors.ExtError) {
    var statusCode = extErr.HTTPCode()

    w.Header().Set("Content-Type", "application/json;charset=utf-8")
    errResp, err := json.Marshal(extErr)
    if err != nil {
        statusCode = http.StatusInternalServerError
    }
    
    w.WriteHeader(statusCode)
    w.Write(errResp)
}
```

Use `ExtErrors` for handling multiple errors
```go
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")    
    
    extErrs := exterrors.NewExtErrors()

    user1, extErr := GetUserByID("user_id_1")
    if extErr != nil {
        extErrs.Add(extErr)
    }

    user2, extErr := GetUserByID("user_id_2")
    if extErr != nil {
        extErrs.Add(extErr)
    }

    SendExtErrors(w, extErrs)
}

func SendExtErrors(w http.ResponseWriter, extErrs exterrors.ExtErrors) {
    var statusCode = extErrs.HTTPCode() // Set the status code according to error type

    w.Header().Set("Content-Type", "application/json;charset=utf-8")
    errResp, err := json.Marshal(extErrs)
    if err != nil {
        statusCode = http.StatusInternalServerError
    }
    
    w.WriteHeader(statusCode)
    w.Write(errResp)
}
```

## Caveats
As `ExtError` requires implementation of standard `error` interface to be compatible with it, be careful, when trying to assign function result `ExtError` to the variable with standard `error` type.
This could cause unpredictable behavior.
