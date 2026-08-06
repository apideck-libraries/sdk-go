# ConflictResponseDetail

Contains parameter or domain specific information related to the error and why it occurred.


## Supported Types

### 

```go
conflictResponseDetail := apierrors.CreateConflictResponseDetailStr(string{/* values here */})
```

### 

```go
conflictResponseDetail := apierrors.CreateConflictResponseDetailMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch conflictResponseDetail.Type {
	case apierrors.ConflictResponseDetailTypeStr:
		// conflictResponseDetail.Str is populated
	case apierrors.ConflictResponseDetailTypeMapOfAny:
		// conflictResponseDetail.MapOfAny is populated
}
```
