import "@typespec/http";

@summary("Company SignUp Validation Error")
model CompanySignUpValidationError {
	name?: string[];
  email?: string[];
  password?: string[];
  finalTaxReturn?: string[];
}
