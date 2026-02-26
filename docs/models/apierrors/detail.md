# Detail

Contains parameter or domain specific information related to the error and why it occurred.


## Supported Types

### 

```go
detail := apierrors.CreateDetailStr(string{/* values here */})
```

### 

```go
detail := apierrors.CreateDetailMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch detail.Type {
	case apierrors.DetailTypeStr:
		// detail.Str is populated
	case apierrors.DetailTypeMapOfAny:
		// detail.MapOfAny is populated
}
```
