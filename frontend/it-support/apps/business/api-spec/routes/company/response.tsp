@summary("Company SignIn Ok Response")
model CompanySignInOkResponse {
  token: string;
}

@summary("Company SignIn Error Response")
model CompanySignInOkResponseHeader {
  "Set-Cookie": string;
}

@summary("Company SignIn BadRequestError Response")
model CompanySignInBadRequestResponse {
  errors: string[];
}
