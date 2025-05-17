import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./request-body.tsp";
import "./response.tsp";

import "../common";

using Http;
using OpenAPI;

@tag("supporters")
@route("/supporters")
namespace RegistrationService.Supporter {
  interface ValidateSignUp {
    @route("/validate-sign-up")
    @operationId("post-supporter-validate-sign-up")
    @summary("Supporter Validate Sign Up")
    @post
    post(
      @header contentType: "multipart/form-data",
      @multipartBody body: SupporterSignUpInput,
    ): SuccessResponse<SupporterSignUpResponse>
      | BadRequestErrorResponse<SupporterSignUpResponse>
      | InternalServerErrorResponse;
  }

  interface SignUp {
    @route("/sign-up")
    @operationId("post-supporter-sign-up")
    @summary("Supporter Sign Up")
    @post
    post(
      @header contentType: "multipart/form-data",
      @multipartBody body: SupporterSignUpInput,
    ): SuccessResponse<SupporterSignUpResponse>
      | BadRequestErrorResponse<SupporterSignUpResponse>
      | InternalServerErrorResponse;
  }
}
