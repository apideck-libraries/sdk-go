# CustomField1Value


## Supported Types

### 

```go
customField1Value := components.CreateCustomField1ValueStr(string{/* values here */})
```

### 

```go
customField1Value := components.CreateCustomField1ValueNumber(float64{/* values here */})
```

### 

```go
customField1Value := components.CreateCustomField1ValueBoolean(bool{/* values here */})
```

### 

```go
customField1Value := components.CreateCustomField1ValueMapOfAny(map[string]any{/* values here */})
```

### 

```go
customField1Value := components.CreateCustomField1ValueArrayOfCustomFieldValue5([]*components.CustomFieldValue5{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customField1Value.Type {
	case components.CustomField1ValueTypeStr:
		// customField1Value.Str is populated
	case components.CustomField1ValueTypeNumber:
		// customField1Value.Number is populated
	case components.CustomField1ValueTypeBoolean:
		// customField1Value.Boolean is populated
	case components.CustomField1ValueTypeMapOfAny:
		// customField1Value.MapOfAny is populated
	case components.CustomField1ValueTypeArrayOfCustomFieldValue5:
		// customField1Value.ArrayOfCustomFieldValue5 is populated
}
```
