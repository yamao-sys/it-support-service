import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

using Http;

@format("binary")
scalar Binary extends string;

@summary("Company SignUp Input")
model CompanySignUpInput {
  name: HttpPart<string>;
  email: HttpPart<string>;
  password: HttpPart<string>;
  finalTaxReturn?: HttpPart<Binary>;
}
