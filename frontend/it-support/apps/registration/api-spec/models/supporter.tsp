import "@typespec/http";

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
