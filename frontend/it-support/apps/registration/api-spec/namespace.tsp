import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

using Http;

@service
@doc("registration APIs")
@OpenAPI.info(#{
  version: "1.0",
  title: "Registration Service",
})
@server("https://localhost:8080")
namespace RegistrationService {};
