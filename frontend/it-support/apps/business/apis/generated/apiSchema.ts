export interface paths {
  "/csrf": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get Csrf */
    get: operations["get-csrf"];
    put?: never;
    post?: never;
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/supporters/signIn": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Supporter SignIn */
    post: operations["post-supporters-sign_in"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/companies/signIn": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** SignIn */
    post: operations["post-companies-sign_in"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
}
export type webhooks = Record<string, never>;
export interface components {
  schemas: never;
  responses: {
    /** @description Csrf response */
    CsrfResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          csrfToken: string;
        };
      };
    };
    /** @description Supporter SignIn Response */
    SupporterSignInOkResponse: {
      headers: {
        "Set-Cookie"?: string;
        [name: string]: unknown;
      };
      content: {
        "application/json": Record<string, never>;
      };
    };
    /** @description Supporter SignIn BadRequest Response */
    SupporterSignInBadRequestResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          errors: string[];
        };
      };
    };
    /** @description Company SignIn Response */
    CompanySignInOkResponse: {
      headers: {
        "Set-Cookie"?: string;
        [name: string]: unknown;
      };
      content: {
        "application/json": Record<string, never>;
      };
    };
    /** @description Company SignIn BadRequest Response */
    CompanySignInBadRequestResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          errors: string[];
        };
      };
    };
    /** @description Internal Server Error Response */
    InternalServerErrorResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          code: number;
          message: string;
        };
      };
    };
  };
  parameters: never;
  requestBodies: {
    /** @description Supporter SignIn  Input */
    SupporterSignInInput: {
      content: {
        "application/json": {
          email: string;
          password: string;
        };
      };
    };
    /** @description Company SignIn  Input */
    CompanySignInInput: {
      content: {
        "application/json": {
          email: string;
          password: string;
        };
      };
    };
  };
  headers: never;
  pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
  "get-csrf": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["CsrfResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-supporters-sign_in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["SupporterSignInInput"];
    responses: {
      200: components["responses"]["SupporterSignInOkResponse"];
      400: components["responses"]["SupporterSignInBadRequestResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-companies-sign_in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["CompanySignInInput"];
    responses: {
      200: components["responses"]["CompanySignInOkResponse"];
      400: components["responses"]["CompanySignInBadRequestResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
}
