# HolderRelationship

The holder's relationship to the account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.HolderRelationshipAuthorizedSigner

// Open enum: custom values can be created with a direct type cast
custom := components.HolderRelationship("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `HolderRelationshipAuthorizedSigner`                     | authorized_signer                                        |
| `HolderRelationshipAuthorizedUser`                       | authorized_user                                          |
| `HolderRelationshipBusiness`                             | business                                                 |
| `HolderRelationshipForBenefitOf`                         | for_benefit_of                                           |
| `HolderRelationshipForBenefitOfPrimary`                  | for_benefit_of_primary                                   |
| `HolderRelationshipForBenefitOfPrimaryJointRestricted`   | for_benefit_of_primary_joint_restricted                  |
| `HolderRelationshipForBenefitOfSecondary`                | for_benefit_of_secondary                                 |
| `HolderRelationshipForBenefitOfSecondaryJointRestricted` | for_benefit_of_secondary_joint_restricted                |
| `HolderRelationshipForBenefitOfSoleOwnerRestricted`      | for_benefit_of_sole_owner_restricted                     |
| `HolderRelationshipPowerOfAttorney`                      | power_of_attorney                                        |
| `HolderRelationshipPrimary`                              | primary                                                  |
| `HolderRelationshipPrimaryBorrower`                      | primary_borrower                                         |
| `HolderRelationshipPrimaryJoint`                         | primary_joint                                            |
| `HolderRelationshipPrimaryJointTenants`                  | primary_joint_tenants                                    |
| `HolderRelationshipSecondary`                            | secondary                                                |
| `HolderRelationshipSecondaryBorrower`                    | secondary_borrower                                       |
| `HolderRelationshipSecondaryJoint`                       | secondary_joint                                          |
| `HolderRelationshipSecondaryJointTenants`                | secondary_joint_tenants                                  |
| `HolderRelationshipSoleOwner`                            | sole_owner                                               |
| `HolderRelationshipTrustee`                              | trustee                                                  |
| `HolderRelationshipUniformTransferToMinor`               | uniform_transfer_to_minor                                |