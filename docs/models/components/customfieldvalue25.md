# CustomFieldValue25


## Supported Types

### 

```go
customFieldValue25 := components.CreateCustomFieldValue25Str(string{/* values here */})
```

### 

```go
customFieldValue25 := components.CreateCustomFieldValue25Number(float64{/* values here */})
```

### 

```go
customFieldValue25 := components.CreateCustomFieldValue25Boolean(bool{/* values here */})
```

### 

```go
customFieldValue25 := components.CreateCustomFieldValue25MapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldValue25.Type {
	case components.CustomFieldValue25TypeStr:
		// customFieldValue25.Str is populated
	case components.CustomFieldValue25TypeNumber:
		// customFieldValue25.Number is populated
	case components.CustomFieldValue25TypeBoolean:
		// customFieldValue25.Boolean is populated
	case components.CustomFieldValue25TypeMapOfAny:
		// customFieldValue25.MapOfAny is populated
}
```
