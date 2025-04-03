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
  "/supporters/validateSignUp": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Validate SignUp */
    post: operations["post-supporter-validate_sign_up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/supporters/signUp": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** SignUp */
    post: operations["post-supporter-sign_up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/companies/validateSignUp": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Validate SignUp */
    post: operations["post-company-validate_sign_up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/companies/signUp": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** SignUp */
    post: operations["post-company-sign_up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
}
export type webhooks = Record<string, never>;
export interface components {
  schemas: {
    /** SuppoterSignUpValidationError */
    SuppoterSignUpValidationError: {
      firstName?: string[];
      lastName?: string[];
      email?: string[];
      password?: string[];
      birthday?: string[];
      frontIdentification?: string[];
      backIdentification?: string[];
    };
    /** CompanySignUpValidationError */
    CompanySignUpValidationError: {
      name?: string[];
      email?: string[];
      password?: string[];
      finalTaxReturn?: string[];
    };
  };
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
    SupporterSignUpResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          code: number;
          errors: components["schemas"]["SuppoterSignUpValidationError"];
        };
      };
    };
    /** @description SignIn Response */
    SupporterSignInOkResponse: {
      headers: {
        "Set-Cookie"?: string;
        [name: string]: unknown;
      };
      content: {
        "application/json": Record<string, never>;
      };
    };
    CompanySignUpResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          code: number;
          errors: components["schemas"]["CompanySignUpValidationError"];
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
          /** Format: int64 */
          code: number;
          message: string;
        };
      };
    };
  };
  parameters: never;
  requestBodies: {
    /** @description Supporter SignUp Iuput */
    SupporterSignUpInput: {
      content: {
        "multipart/form-data": {
          firstName: string;
          lastName: string;
          email: string;
          password: string;
          /** Format: date */
          birthday?: string;
          /** Format: binary */
          frontIdentification?: Blob;
          /** Format: binary */
          backIdentification?: Blob;
        };
      };
    };
    /** @description Company SignUp Iuput */
    CompanySignUpInput: {
      content: {
        "multipart/form-data": {
          name: string;
          email: string;
          password: string;
          /** Format: binary */
          finalTaxReturn?: Blob;
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
  "post-supporter-validate_sign_up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["SupporterSignUpInput"];
    responses: {
      200: components["responses"]["SupporterSignUpResponse"];
      400: components["responses"]["SupporterSignUpResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-supporter-sign_up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["SupporterSignUpInput"];
    responses: {
      200: components["responses"]["SupporterSignUpResponse"];
      400: components["responses"]["SupporterSignUpResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-company-validate_sign_up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["CompanySignUpInput"];
    responses: {
      200: components["responses"]["CompanySignUpResponse"];
      400: components["responses"]["CompanySignUpResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-company-sign_up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["CompanySignUpInput"];
    responses: {
      200: components["responses"]["CompanySignUpResponse"];
      400: components["responses"]["CompanySignUpResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
}
