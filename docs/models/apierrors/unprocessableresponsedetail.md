# UnprocessableResponseDetail

Contains parameter or domain specific information related to the error and why it occurred.


## Supported Types

### 

```go
unprocessableResponseDetail := apierrors.CreateUnprocessableResponseDetailStr(string{/* values here */})
```

### 

```go
unprocessableResponseDetail := apierrors.CreateUnprocessableResponseDetailMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch unprocessableResponseDetail.Type {
	case apierrors.UnprocessableResponseDetailTypeStr:
		// unprocessableResponseDetail.Str is populated
	case apierrors.UnprocessableResponseDetailTypeMapOfAny:
		// unprocessableResponseDetail.MapOfAny is populated
}
```
