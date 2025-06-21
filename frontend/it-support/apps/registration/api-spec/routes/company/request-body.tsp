import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "../common";

using Http;

@summary("Company SignUp Input")
model CompanySignUpInput {
  name: HttpPart<string>;
  email: HttpPart<string>;
  password: HttpPart<string>;
  finalTaxReturn?: HttpPart<Binary>;
}
