export interface paths {
  "/companies/sign-up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Company Sign Up */
    post: operations["post-company-sign-up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/companies/validate-sign-up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Company Validate Sign Up */
    post: operations["post-company-validate-sign-up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
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
  "/supporters/sign-up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Supporter Sign Up */
    post: operations["post-supporter-sign-up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/supporters/validate-sign-up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Supporter Validate Sign Up */
    post: operations["post-supporter-validate-sign-up"];
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
    CompanySignUpInput: {
      name: string;
      email: string;
      password: string;
      /** Format: binary */
      finalTaxReturn?: Blob;
    };
    /** Company SignUp Response */
    CompanySignUpResponse: {
      code: number;
      errors: components["schemas"]["CompanySignUpValidationError"];
    };
    /** Company SignUp Validation Error */
    CompanySignUpValidationError: {
      name?: string[];
      email?: string[];
      password?: string[];
      finalTaxReturn?: string[];
    };
    /** CsrfResponse */
    CsrfResponse: {
      csrfToken: string;
    };
    SupporterSignUpInput: {
      firstName: string;
      lastName: string;
      email: string;
      password: string;
      /** Format: date */
      birthday?: Date;
      /** Format: binary */
      frontIdentification?: Blob;
      /** Format: binary */
      backIdentification?: Blob;
    };
    /** Supporter SignUp Response */
    SupporterSignUpResponse: {
      code: number;
      errors: components["schemas"]["SupporterSignUpValidationError"];
    };
    /** Supporter SignUp Validation Error */
    SupporterSignUpValidationError: {
      firstName?: string[];
      lastName?: string[];
      email?: string[];
      password?: string[];
      birthday?: string[];
      frontIdentification?: string[];
      backIdentification?: string[];
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
  "post-company-sign-up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody: {
      content: {
        "multipart/form-data": components["schemas"]["CompanySignUpInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["CompanySignUpResponse"];
        };
      };
      /** @description The server could not understand the request due to invalid syntax. */
      400: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["CompanySignUpResponse"];
        };
      };
      /** @description Server error */
      500: {
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
  };
  "post-company-validate-sign-up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody: {
      content: {
        "multipart/form-data": components["schemas"]["CompanySignUpInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["CompanySignUpResponse"];
        };
      };
      /** @description The server could not understand the request due to invalid syntax. */
      400: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["CompanySignUpResponse"];
        };
      };
      /** @description Server error */
      500: {
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
  };
  "get-csrf": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["CsrfResponse"];
        };
      };
      /** @description Server error */
      500: {
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
  };
  "post-supporter-sign-up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody: {
      content: {
        "multipart/form-data": components["schemas"]["SupporterSignUpInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["SupporterSignUpResponse"];
        };
      };
      /** @description The server could not understand the request due to invalid syntax. */
      400: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["SupporterSignUpResponse"];
        };
      };
      /** @description Server error */
      500: {
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
  };
  "post-supporter-validate-sign-up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody: {
      content: {
        "multipart/form-data": components["schemas"]["SupporterSignUpInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["SupporterSignUpResponse"];
        };
      };
      /** @description The server could not understand the request due to invalid syntax. */
      400: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["SupporterSignUpResponse"];
        };
      };
      /** @description Server error */
      500: {
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
  };
}
