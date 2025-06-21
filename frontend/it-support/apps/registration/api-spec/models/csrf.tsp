import "@typespec/http";

@summary("CsrfResponse")
model CsrfResponse {
  csrfToken: string;
}
