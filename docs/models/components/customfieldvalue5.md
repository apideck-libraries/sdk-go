# CustomFieldValue5


## Supported Types

### 

```go
customFieldValue5 := components.CreateCustomFieldValue5Str(string{/* values here */})
```

### 

```go
customFieldValue5 := components.CreateCustomFieldValue5Number(float64{/* values here */})
```

### 

```go
customFieldValue5 := components.CreateCustomFieldValue5Boolean(bool{/* values here */})
```

### 

```go
customFieldValue5 := components.CreateCustomFieldValue5MapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldValue5.Type {
	case components.CustomFieldValue5TypeStr:
		// customFieldValue5.Str is populated
	case components.CustomFieldValue5TypeNumber:
		// customFieldValue5.Number is populated
	case components.CustomFieldValue5TypeBoolean:
		// customFieldValue5.Boolean is populated
	case components.CustomFieldValue5TypeMapOfAny:
		// customFieldValue5.MapOfAny is populated
}
```
