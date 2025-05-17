import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./request-body.tsp";
import "./response.tsp";

import "../common";

using Http;
using OpenAPI;

@tag("plans")
@route("/plans")
@useAuth([SecuritySchema])
namespace BusinessService.Plan {
  interface Root {
    @operationId("post-plans")
    @summary("Create Plan")
    @post
    post(
      @body body: PlanStoreInput,
    ): SuccessResponse<PlanStoreResponse>
      | InternalServerErrorResponse;
  }
}
