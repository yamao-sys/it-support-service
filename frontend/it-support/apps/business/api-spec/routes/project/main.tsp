import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./request-body.tsp";
import "./response.tsp";

import "../common";

using Http;
using OpenAPI;

@tag("projects")
@route("/projects")
@useAuth([SecuritySchema])
namespace BusinessService.Project {
  interface Root {
    @operationId("post-project")
    @summary("Create Project")
    @post
    post(
      @body body: ProjectStoreInput,
    ): SuccessResponse<ProjectStoreResponse>
      | InternalServerErrorResponse;
  }
}
