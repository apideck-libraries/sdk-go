# Detail

Contains parameter or domain specific information related to the error and why it occurred.


## Supported Types

### 

```go
detail := components.CreateDetailStr(string{/* values here */})
```

### 

```go
detail := components.CreateDetailMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch detail.Type {
	case components.DetailTypeStr:
		// detail.Str is populated
	case components.DetailTypeMapOfAny:
		// detail.MapOfAny is populated
}
```
