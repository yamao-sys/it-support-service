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
namespace BusinessService.Supporter {
  interface SignIn {
    @route("/sign-in")
    @operationId("post-supporter-sign-in")
    @summary("Supporter Sign In")
    @post
    post(
      @body body: SupporterSignInInput,
    ): SuccessResponseWithHeader<SupporterSignInOkResponse>
      | BadRequestErrorResponse<SupporterSignInBadRequestResponse>
      | InternalServerErrorResponse;
  }
}
