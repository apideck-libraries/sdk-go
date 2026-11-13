# Error

Why the operation is unsupported. Absent when there is no reason to give (the connector does not implement the operation). A string when the connector states a reason. An object `{ reason, error }` when a batch request is refused: `reason` is a stable code, `error` a human-readable explanation.


## Supported Types

### 

```go
error := apierrors.CreateErrorStr(string{/* values here */})
```

### Error2

```go
error := apierrors.CreateErrorError2(apierrors.Error2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch error.Type {
	case apierrors.ErrorTypeStr:
		// error.Str is populated
	case apierrors.ErrorTypeError2:
		// error.Error2 is populated
}
```
