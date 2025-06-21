import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "../../models/plan.tsp";
import "../common";

using Http;
using OpenAPI;

@tag("plans")
@route("/plans")
@useAuth([SecuritySchema])
namespace BusinessService.Plan {
  interface Root {
    @operationId("post-plan")
    @summary("Create Plan")
    @post
    post(
      @body body: PlanStoreInput,
    ): SuccessResponse<PlanStoreResponse>
      | InternalServerErrorResponse;
  }
}
