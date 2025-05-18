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
    @operationId("get-projects")
    @summary("Create Project")
    @get
    get(
      @query pageToken?: string,
    ): SuccessResponse<ProjectsListResponse>
      | InternalServerErrorResponse;

    @operationId("post-project")
    @summary("Create Project")
    @post
    post(
      @body body: ProjectStoreInput,
    ): SuccessResponse<ProjectStoreResponse>
      | InternalServerErrorResponse;
  }

  @route("/{id}")
  interface Id {
    @operationId("get-project")
    @summary("Get Project")
    @get
    get(
      @path id: integer,
    ): SuccessResponse<ProjectResponse>
      | NotFoundErrorResponse
      | InternalServerErrorResponse;

    @operationId("put-project")
    @summary("Update Project")
    @put
    put(
      @path id: integer,
      @body body: ProjectStoreInput,
    ): SuccessResponse<ProjectStoreResponse>
      | NotFoundErrorResponse
      | InternalServerErrorResponse;
  }
}
