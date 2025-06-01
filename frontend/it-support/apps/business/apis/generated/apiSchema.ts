export interface paths {
  "/companies/sign-in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Company Sign In */
    post: operations["post-company-sign-in"];
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
  "/plans": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Create Plan */
    post: operations["post-plan"];
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
    /** Create Project */
    get: operations["get-projects"];
    put?: never;
    /** Create Project */
    post: operations["post-project"];
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
    /** Get Project */
    get: operations["get-project"];
    /** Update Project */
    put: operations["put-project"];
    post?: never;
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/supporters/sign-in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** Supporter Sign In */
    post: operations["post-supporter-sign-in"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/to-projects": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get Projects for Supporters */
    get: operations["get-to-projects"];
    put?: never;
    post?: never;
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/to-projects/{id}": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get Project for Supporters */
    get: operations["get-to-project"];
    put?: never;
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
    /** Company SignIn BadRequestError Response */
    CompanySignInBadRequestResponse: {
      errors: string[];
    };
    /** Company SignIn Input */
    CompanySignInInput: {
      email: string;
      password: string;
    };
    /** Company SignIn Ok Response */
    CompanySignInOkResponse: Record<string, never>;
    /** CsrfResponse */
    CsrfResponse: {
      csrfToken: string;
    };
    /** Plan */
    Plan: {
      id: number;
      projectId: number;
      title: string;
      description: string;
      /** Format: date */
      startDate: Date;
      /** Format: date */
      endDate: Date;
      unitPrice: number;
      /** Format: date-time */
      createdAt: string;
    };
    /** Plan Store Input */
    PlanStoreInput: {
      projectId: number;
      title: string;
      description: string;
      /** Format: date */
      startDate: Date;
      /** Format: date */
      endDate: Date;
      unitPrice: number;
    };
    /** Plan Store Response */
    PlanStoreResponse: {
      plan: components["schemas"]["Plan"];
      errors: components["schemas"]["PlanValidationError"];
    };
    /** Plan Validation Error */
    PlanValidationError: {
      title?: string[];
      description?: string[];
      startDate?: string[];
      endDate?: string[];
      unitPrice?: string[];
    };
    /** Project */
    Project: {
      id: number;
      title: string;
      description: string;
      /** Format: date */
      startDate: Date;
      /** Format: date */
      endDate: Date;
      minBudget?: number;
      maxBudget?: number;
      isActive: boolean;
      /** Format: date-time */
      createdAt: string;
    };
    /** Project Response */
    ProjectResponse: {
      project: components["schemas"]["Project"];
    };
    /** Project Store Input */
    ProjectStoreInput: {
      title: string;
      description: string;
      /** Format: date */
      startDate?: Date;
      /** Format: date */
      endDate?: Date;
      minBudget?: number;
      maxBudget?: number;
      isActive: boolean;
    };
    /** Project Store Response */
    ProjectStoreResponse: {
      project: components["schemas"]["Project"];
      errors: components["schemas"]["ProjectValidationError"];
    };
    /** Project Validation Error */
    ProjectValidationError: {
      title?: string[];
      description?: string[];
      startDate?: string[];
      endDate?: string[];
      minBudget?: string[];
      maxBudget?: string[];
      isActive?: string[];
    };
    /** Projects List Response */
    ProjectsListResponse: {
      projects: components["schemas"]["Project"][];
      nextPageToken: string;
    };
    /** Supporter SignIn BadRequestError Response */
    SupporterSignInBadRequestResponse: {
      errors: string[];
    };
    /** Supporter SignIn Input */
    SupporterSignInInput: {
      email: string;
      password: string;
    };
    /** Supporter SignIn Ok Response */
    SupporterSignInOkResponse: Record<string, never>;
    /** Project for Supporters */
    ToProject: {
      id: number;
      title: string;
      description: string;
      /** Format: date */
      startDate: Date;
      /** Format: date */
      endDate: Date;
      minBudget?: number;
      maxBudget?: number;
    };
    /** Project Response for Supporters */
    ToProjectResponse: {
      project: components["schemas"]["ToProject"];
    };
    /** Projects List Response for Supporters */
    ToProjectsListResponse: {
      projects: components["schemas"]["ToProject"][];
      nextPageToken: string;
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
  "post-company-sign-in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["CompanySignInInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          "Set-Cookie": string;
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["CompanySignInOkResponse"];
        };
      };
      /** @description The server could not understand the request due to invalid syntax. */
      400: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["CompanySignInBadRequestResponse"];
        };
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
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
        content?: never;
      };
    };
  };
  "post-plan": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["PlanStoreInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["PlanStoreResponse"];
        };
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
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
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["ProjectsListResponse"];
        };
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
    };
  };
  "post-project": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["ProjectStoreInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["ProjectStoreResponse"];
        };
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
    };
  };
  "get-project": {
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
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["ProjectResponse"];
        };
      };
      /** @description The server cannot find the requested resource. */
      404: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
    };
  };
  "put-project": {
    parameters: {
      query?: never;
      header?: never;
      path: {
        id: number;
      };
      cookie?: never;
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["ProjectStoreInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["ProjectStoreResponse"];
        };
      };
      /** @description The server cannot find the requested resource. */
      404: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
    };
  };
  "post-supporter-sign-in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["SupporterSignInInput"];
      };
    };
    responses: {
      /** @description The request has succeeded. */
      200: {
        headers: {
          "Set-Cookie": string;
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["SupporterSignInOkResponse"];
        };
      };
      /** @description The server could not understand the request due to invalid syntax. */
      400: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["SupporterSignInBadRequestResponse"];
        };
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
    };
  };
  "get-to-projects": {
    parameters: {
      query?: {
        pageToken?: string;
        startDate?: Date;
        endDate?: Date;
      };
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
          "application/json": components["schemas"]["ToProjectsListResponse"];
        };
      };
      /** @description Access is forbidden. */
      403: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
    };
  };
  "get-to-project": {
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
      /** @description The request has succeeded. */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": components["schemas"]["ToProjectResponse"];
        };
      };
      /** @description Access is forbidden. */
      403: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
      /** @description The server cannot find the requested resource. */
      404: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
      /** @description Server error */
      500: {
        headers: {
          [name: string]: unknown;
        };
        content?: never;
      };
    };
  };
}
