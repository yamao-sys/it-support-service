@summary("Supporter SignIn Ok Response")
model SupporterSignInOkResponse {}

@summary("Supporter SignIn Error Response")
model SupporterSignInOkResponseHeader {
  "Set-Cookie": string;
}

@summary("Supporter SignIn BadRequestError Response")
model SupporterSignInBadRequestResponse {
  errors: string[];
}
