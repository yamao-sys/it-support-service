import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "../common/request.tsp";

using Http;

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
