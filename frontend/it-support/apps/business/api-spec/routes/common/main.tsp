import "@typespec/http";

using Http;

alias SecuritySchema = ApiKeyAuth<ApiKeyLocation.cookie, "token">;

@format("binary")
scalar Binary extends string;

model SuccessResponse<T> {
  @statusCode status: 200;
  @body body: T
}

model SuccessResponseWithHeader<T> {
  @statusCode status: 200;
  @body body: T;
  @header("Set-Cookie") token: string;
}

model BadRequestErrorResponse<T> {
  @statusCode status: 400;
  @body body: T
}

model ForBiddenErrorResponse {
  @statusCode status: 403;
}

model NotFoundErrorResponse {
  @statusCode status: 404;
}

model InternalServerErrorResponse {
  @statusCode status: 500;
}
