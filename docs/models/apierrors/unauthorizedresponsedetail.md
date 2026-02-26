# UnauthorizedResponseDetail

Contains parameter or domain specific information related to the error and why it occurred.


## Supported Types

### 

```go
unauthorizedResponseDetail := apierrors.CreateUnauthorizedResponseDetailStr(string{/* values here */})
```

### Two

```go
unauthorizedResponseDetail := apierrors.CreateUnauthorizedResponseDetailTwo(apierrors.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch unauthorizedResponseDetail.Type {
	case apierrors.UnauthorizedResponseDetailTypeStr:
		// unauthorizedResponseDetail.Str is populated
	case apierrors.UnauthorizedResponseDetailTypeTwo:
		// unauthorizedResponseDetail.Two is populated
}
```
