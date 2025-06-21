import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "../routes/common";

using Http;

@summary("Supporter SignUp Validation Error")
model SupporterSignUpValidationError {
  firstName?: string[];
  lastName?: string[];
  email?: string[];
  password?: string[];
  birthday?: string[];
  frontIdentification?: string[];
  backIdentification?: string[];
}

@summary("Supporter SignUp Input")
model SupporterSignUpInput {
  firstName: HttpPart<string>;
  lastName: HttpPart<string>;
  email: HttpPart<string>;
  password: HttpPart<string>;
  birthday?: HttpPart<plainDate>;
  frontIdentification?: HttpPart<Binary>;
  backIdentification?: HttpPart<Binary>;
}

@summary("Supporter SignUp Response")
model SupporterSignUpResponse {
  code: integer;
  errors: SupporterSignUpValidationError;
}
