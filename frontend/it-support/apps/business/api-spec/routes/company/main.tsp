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
namespace BusinessService.Company {
  interface SignIn {
    @route("/sign-in")
    @operationId("post-company-sign-in")
    @summary("Company Sign In")
    @post
    post(
      @body body: CompanySignInInput,
    ): SuccessResponseWithHeader<CompanySignInOkResponse>
      | BadRequestErrorResponse<CompanySignInBadRequestResponse>
      | InternalServerErrorResponse;
  }
}
