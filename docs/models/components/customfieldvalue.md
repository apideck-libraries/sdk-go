# CustomFieldValue


## Supported Types

### 

```go
customFieldValue := components.CreateCustomFieldValueStr(string{/* values here */})
```

### 

```go
customFieldValue := components.CreateCustomFieldValueNumber(float64{/* values here */})
```

### 

```go
customFieldValue := components.CreateCustomFieldValueBoolean(bool{/* values here */})
```

### 

```go
customFieldValue := components.CreateCustomFieldValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
customFieldValue := components.CreateCustomFieldValueArrayOfCustomFieldValue25([]*components.CustomFieldValue25{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldValue.Type {
	case components.CustomFieldValueTypeStr:
		// customFieldValue.Str is populated
	case components.CustomFieldValueTypeNumber:
		// customFieldValue.Number is populated
	case components.CustomFieldValueTypeBoolean:
		// customFieldValue.Boolean is populated
	case components.CustomFieldValueTypeMapOfAny:
		// customFieldValue.MapOfAny is populated
	case components.CustomFieldValueTypeArrayOfCustomFieldValue25:
		// customFieldValue.ArrayOfCustomFieldValue25 is populated
}
```
