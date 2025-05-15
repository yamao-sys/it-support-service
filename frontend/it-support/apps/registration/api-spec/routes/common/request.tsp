import "@typespec/http";
import "@typespec/openapi";

using Http;

alias SecuritySchema = ApiKeyAuth<ApiKeyLocation.cookie, "token">;
