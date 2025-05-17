import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./request-body.tsp";
import "./response.tsp";

import "../common";

using Http;
using OpenAPI;

@tag("companies")
@route("/companies")
@useAuth([SecuritySchema])
namespace RegistrationService.Company {
  interface ValidateSignUp {
    @route("/validateSignUp")
    @operationId("post-company-validate_sign_up")
    @summary("Company Validate Sign Up")
    @post
    post(
      @header contentType: "multipart/form-data",
      @multipartBody body: CompanySignUpInput,
    ): SuccessResponse<CompanySignUpResponse>
      | BadRequestErrorResponse<CompanySignUpResponse>
      | InternalServerErrorResponse;
  }

  interface SignUp {
    @route("/signUp")
    @operationId("post-company-sign_up")
    @summary("Company Sign Up")
    @post
    post(
      @header contentType: "multipart/form-data",
      @multipartBody body: CompanySignUpInput,
    ): SuccessResponse<CompanySignUpResponse>
      | BadRequestErrorResponse<CompanySignUpResponse>
      | InternalServerErrorResponse;
  }
}
