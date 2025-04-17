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
  "/projects": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Project List */
    get: operations["get-projects"];
    put?: never;
    /** Project Create */
    post: operations["post-projects"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/projects/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Project Show */
    get: operations["get-projects-id"];
    /** Project Update */
    put: operations["put-projects-id"];
    post?: never;
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
    /**
     * Project Model
     * @description Project
     */
    Project: {
      id?: string;
      title?: string;
      description?: string;
      /** Format: date */
      start_date?: Date;
      /** Format: date */
      end_date?: Date;
      min_budget?: number;
      max_budget?: number;
      isActive?: boolean;
      /** Format: date-time */
      created_at?: string;
    };
    /** ProjectValidationErrors */
    ProjectValidationError: {
      title?: string[];
      description?: string[];
      startDate?: string[];
      endDate?: string[];
      minBudget?: string[];
      maxBudget?: string[];
      isActive?: string[];
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
    /** @description Project Store Response */
    ProjectStoreResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          project: components["schemas"]["Project"];
          errors: components["schemas"]["ProjectValidationError"];
        };
      };
    };
    /** @description Projects List Response */
    ProjectsListResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          projects: components["schemas"]["Project"][];
          nextPageToken: string;
        };
      };
    };
    /** @description Project Response */
    ProjectResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          project: components["schemas"]["Project"];
        };
      };
    };
    /** @description Not Found Error Response */
    NotFoundErrorResponse: {
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
    /** @description Project Store Inputs */
    ProjectStoreInput: {
      content: {
        "application/json": {
          title?: string;
          description?: string;
          /** Format: date */
          startDate?: Date;
          /** Format: date */
          endDate?: Date;
          minBudget?: number;
          maxBudget?: number;
          isActive?: boolean;
        };
      };
    };
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
  "get-projects": {
    parameters: {
      query?: {
        pageToken?: string;
      };
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["ProjectsListResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-projects": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["ProjectStoreInput"];
    responses: {
      200: components["responses"]["ProjectStoreResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "get-projects-id": {
    parameters: {
      query?: never;
      header?: never;
      path: {
        id: number;
      };
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["ProjectResponse"];
      404: components["responses"]["NotFoundErrorResponse"];
    };
  };
  "put-projects-id": {
    parameters: {
      query?: never;
      header?: never;
      path: {
        id: number;
      };
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["ProjectStoreInput"];
    responses: {
      200: components["responses"]["ProjectStoreResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
}
