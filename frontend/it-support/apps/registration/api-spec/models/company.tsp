import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "../routes/common";

using Http;

@summary("Company SignUp Validation Error")
model CompanySignUpValidationError {
  name?: string[];
  email?: string[];
  password?: string[];
  finalTaxReturn?: string[];
}

@summary("Company SignUp Input")
model CompanySignUpInput {
  name: HttpPart<string>;
  email: HttpPart<string>;
  password: HttpPart<string>;
  finalTaxReturn?: HttpPart<Binary>;
}

@summary("Company SignUp Response")
model CompanySignUpResponse {
  code: integer;
  errors: CompanySignUpValidationError;
}
