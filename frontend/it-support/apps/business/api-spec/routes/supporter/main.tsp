import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "../../models/supporter.tsp";
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
