# CustomField


## Supported Types

### CustomField1

```go
customField := components.CreateCustomFieldCustomField1(components.CustomField1{/* values here */})
```

### CustomField2

```go
customField := components.CreateCustomFieldCustomField2(components.CustomField2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customField.Type {
	case components.CustomFieldTypeCustomField1:
		// customField.CustomField1 is populated
	case components.CustomFieldTypeCustomField2:
		// customField.CustomField2 is populated
}
```
