# UnauthorizedDetail

Contains parameter or domain specific information related to the error and why it occurred.


## Supported Types

### 

```go
unauthorizedDetail := apierrors.CreateUnauthorizedDetailStr(string{/* values here */})
```

### Detail2

```go
unauthorizedDetail := apierrors.CreateUnauthorizedDetailDetail2(apierrors.Detail2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch unauthorizedDetail.Type {
	case apierrors.UnauthorizedDetailTypeStr:
		// unauthorizedDetail.Str is populated
	case apierrors.UnauthorizedDetailTypeDetail2:
		// unauthorizedDetail.Detail2 is populated
}
```
