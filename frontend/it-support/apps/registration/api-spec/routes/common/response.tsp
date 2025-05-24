import "@typespec/http";
import "@typespec/openapi";

using Http;

model SuccessResponse<T> {
  @statusCode status: 200;
  @body body: T
}

model BadRequestErrorResponse<T> {
  @statusCode status: 400;
  @body body: T
}

model NotFoundErrorResponse {
  @statusCode status: 404;
}

model InternalServerErrorResponse {
  @statusCode status: 500;
}
