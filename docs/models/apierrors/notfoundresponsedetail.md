# NotFoundResponseDetail

Contains parameter or domain specific information related to the error and why it occurred.


## Supported Types

### 

```go
notFoundResponseDetail := apierrors.CreateNotFoundResponseDetailStr(string{/* values here */})
```

### 

```go
notFoundResponseDetail := apierrors.CreateNotFoundResponseDetailMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch notFoundResponseDetail.Type {
	case apierrors.NotFoundResponseDetailTypeStr:
		// notFoundResponseDetail.Str is populated
	case apierrors.NotFoundResponseDetailTypeMapOfAny:
		// notFoundResponseDetail.MapOfAny is populated
}
```
