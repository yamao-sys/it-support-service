import "../../models/company.tsp";

@summary("Company SignUp Response")
model CompanySignUpResponse {
	code: integer;
	errors: CompanySignUpValidationError;
}
