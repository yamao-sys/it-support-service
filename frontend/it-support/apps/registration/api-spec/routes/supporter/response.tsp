import "../../models/supporter.tsp";

@summary("Supporter SignUp Response")
model SupporterSignUpResponse {
  code: integer;
  errors: SupporterSignUpValidationError;
}
