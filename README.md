## Golp Validator

**Golp Validator** is a lightweight, chainable validation library for Go, inspired by the simplicity and expressiveness of [Joi](https://joi.dev/).

Designed for struct field validation, it provides a clean, fluent API for defining rules like `Required()`, `MinLength()`, `Email()`, and custom validators — all without heavy dependencies.

### Why Golp Validator?

* ✅ Inspired by Joi, built for Go
* ✅ Lightweight, minimal runtime overhead
* ✅ Focused on clear error reporting
* ✅ Supports chaining with custom error messages
* ✅ No reflection-based magic — simple, explicit validation

### Example Usage

```go
err := validator.ValidateStruct(&user,
	validator.Field(&user.Email, validator.Required(), validator.Email()),
	validator.Field(&user.Age, validator.Required(), validator.Min(18)),
)
if err != nil {
	// handle validation error
}
```

#### With struct

```go
type RegisterRequest struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
}

func (r RegisterRequest) Validate() error {
	return validator.ValidateStruct(&r,
		validator.Field(&r.Username,
			validator.Required().Error("Username is required"),
			validator.MinLength(3).Error("At least 3 characters"),
			validator.MaxLength(50).Error("At most 50 characters"),
		),
		validator.Field(&r.Password,
			validator.Required().Error("Password is required"),
			validator.MinLength(6).Error("At least 6 characters"),
			validator.MaxLength(100).Error("At most 100 characters"),
		),
	)
}
```