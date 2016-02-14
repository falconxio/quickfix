package commissiondata

//Component is a fix50sp1 CommissionData Component
type Component struct {
	//Commission is a non-required field for CommissionData.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for CommissionData.
	CommType *string `fix:"13"`
	//CommCurrency is a non-required field for CommissionData.
	CommCurrency *string `fix:"479"`
	//FundRenewWaiv is a non-required field for CommissionData.
	FundRenewWaiv *string `fix:"497"`
}

func New() *Component { return new(Component) }