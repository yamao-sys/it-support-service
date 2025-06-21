import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "../../models/csrf.tsp";
import "../common";

using Http;
using OpenAPI;

@tag("csrf")
@route("/csrf")
namespace BusinessService.Csrf {
  interface Root {
    @operationId("get-csrf")
    @summary("Get Csrf")
    @get
    get(): SuccessResponse<CsrfResponse>
      | InternalServerErrorResponse;
  }
}
