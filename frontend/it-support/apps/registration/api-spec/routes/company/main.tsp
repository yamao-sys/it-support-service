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
namespace RegistrationService.Company {
  interface ValidateSignUp {
    @route("/validate-sign-up")
    @operationId("post-company-validate-sign-up")
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
    @route("/sign-up")
    @operationId("post-company-sign-up")
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
