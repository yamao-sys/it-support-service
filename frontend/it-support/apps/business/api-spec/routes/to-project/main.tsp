import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./response.tsp";

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
}
