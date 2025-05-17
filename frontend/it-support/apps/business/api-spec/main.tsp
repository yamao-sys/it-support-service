import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

using Http;

@service
@doc("business APIs")
@OpenAPI.info(#{
  version: "1.0",
  title: "Business Service",
})
@server("https://localhost:8080")
namespace SampleService {};
