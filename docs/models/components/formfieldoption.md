# FormFieldOption


## Supported Types

### SimpleFormFieldOption

```go
formFieldOption := components.CreateFormFieldOptionSimple(components.SimpleFormFieldOption{/* values here */})
```

### FormFieldOptionGroup

```go
formFieldOption := components.CreateFormFieldOptionGroup(components.FormFieldOptionGroup{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch formFieldOption.Type {
	case components.FormFieldOptionTypeSimple:
		// formFieldOption.SimpleFormFieldOption is populated
	case components.FormFieldOptionTypeGroup:
		// formFieldOption.FormFieldOptionGroup is populated
}
```
