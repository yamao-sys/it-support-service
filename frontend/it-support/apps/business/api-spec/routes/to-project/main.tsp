import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "../../models/to-project.tsp";
import "../../models/plan.tsp";
import "../common";

using Http;
using OpenAPI;

@tag("to-projects")
@route("/to-projects")
@useAuth([SecuritySchema])
namespace BusinessService.ToProject {
  interface Root {
    @operationId("get-to-projects")
    @summary("Get Projects for Supporters")
    @get
    get(
      @query(#{ explode: true }) pageToken?: string,
      @query(#{ explode: true }) startDate?: plainDate,
      @query(#{ explode: true }) endDate?: plainDate,
    ): SuccessResponse<ToProjectsListResponse>
      | ForBiddenErrorResponse
      | InternalServerErrorResponse;
  }

  @route("/{id}")
  interface Id {
    @operationId("get-to-project")
    @summary("Get Project for Supporters")
    @get
    get(
      @path id: integer,
    ): SuccessResponse<ToProjectResponse>
      | ForBiddenErrorResponse
      | NotFoundErrorResponse
      | InternalServerErrorResponse;
  }

  @route("/{id}/plans")
  interface IdPlans {
    @operationId("post-to-project-plan")
    @summary("Create Plan for Project")
    @post
    post(
      @path id: integer,
      @body body: PlanStoreWithStepsInput,
    ): SuccessResponse<PlanWithStepsStoreResponse>
      | BadRequestErrorResponse<PlanWithStepsValidationError>
      | ForBiddenErrorResponse
      | NotFoundErrorResponse
      | InternalServerErrorResponse;
  }
}
