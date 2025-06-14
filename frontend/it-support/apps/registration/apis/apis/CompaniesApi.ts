/* tslint:disable */
/* eslint-disable */
/**
 * Registration Service
 * registration APIs
 *
 * The version of the OpenAPI document: 1.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import * as runtime from "../runtime";
import type { CompanySignUpResponse } from "../models/index";
import { CompanySignUpResponseFromJSON, CompanySignUpResponseToJSON } from "../models/index";

export interface PostCompanySignUpRequest {
  name: string;
  email: string;
  password: string;
  finalTaxReturn?: Blob;
}

export interface PostCompanyValidateSignUpRequest {
  name: string;
  email: string;
  password: string;
  finalTaxReturn?: Blob;
}

/**
 *
 */
export class CompaniesApi extends runtime.BaseAPI {
  /**
   * Company Sign Up
   */
  async postCompanySignUpRaw(
    requestParameters: PostCompanySignUpRequest,
    initOverrides?: RequestInit | runtime.InitOverrideFunction,
  ): Promise<runtime.ApiResponse<CompanySignUpResponse>> {
    if (requestParameters["name"] == null) {
      throw new runtime.RequiredError(
        "name",
        'Required parameter "name" was null or undefined when calling postCompanySignUp().',
      );
    }

    if (requestParameters["email"] == null) {
      throw new runtime.RequiredError(
        "email",
        'Required parameter "email" was null or undefined when calling postCompanySignUp().',
      );
    }

    if (requestParameters["password"] == null) {
      throw new runtime.RequiredError(
        "password",
        'Required parameter "password" was null or undefined when calling postCompanySignUp().',
      );
    }

    const queryParameters: any = {};

    const headerParameters: runtime.HTTPHeaders = {};

    const consumes: runtime.Consume[] = [{ contentType: "multipart/form-data" }];
    // @ts-ignore: canConsumeForm may be unused
    const canConsumeForm = runtime.canConsumeForm(consumes);

    let formParams: { append(param: string, value: any): any };
    let useForm = false;
    // use FormData to transmit files using content-type "multipart/form-data"
    useForm = canConsumeForm;
    if (useForm) {
      formParams = new FormData();
    } else {
      formParams = new URLSearchParams();
    }

    if (requestParameters["name"] != null) {
      formParams.append("name", requestParameters["name"] as any);
    }

    if (requestParameters["email"] != null) {
      formParams.append("email", requestParameters["email"] as any);
    }

    if (requestParameters["password"] != null) {
      formParams.append("password", requestParameters["password"] as any);
    }

    if (requestParameters["finalTaxReturn"] != null) {
      formParams.append("finalTaxReturn", requestParameters["finalTaxReturn"] as any);
    }

    const response = await this.request(
      {
        path: `/companies/sign-up`,
        method: "POST",
        headers: headerParameters,
        query: queryParameters,
        body: formParams,
      },
      initOverrides,
    );

    return new runtime.JSONApiResponse(response, (jsonValue) =>
      CompanySignUpResponseFromJSON(jsonValue),
    );
  }

  /**
   * Company Sign Up
   */
  async postCompanySignUp(
    requestParameters: PostCompanySignUpRequest,
    initOverrides?: RequestInit | runtime.InitOverrideFunction,
  ): Promise<CompanySignUpResponse> {
    const response = await this.postCompanySignUpRaw(requestParameters, initOverrides);
    return await response.value();
  }

  /**
   * Company Validate Sign Up
   */
  async postCompanyValidateSignUpRaw(
    requestParameters: PostCompanyValidateSignUpRequest,
    initOverrides?: RequestInit | runtime.InitOverrideFunction,
  ): Promise<runtime.ApiResponse<CompanySignUpResponse>> {
    if (requestParameters["name"] == null) {
      throw new runtime.RequiredError(
        "name",
        'Required parameter "name" was null or undefined when calling postCompanyValidateSignUp().',
      );
    }

    if (requestParameters["email"] == null) {
      throw new runtime.RequiredError(
        "email",
        'Required parameter "email" was null or undefined when calling postCompanyValidateSignUp().',
      );
    }

    if (requestParameters["password"] == null) {
      throw new runtime.RequiredError(
        "password",
        'Required parameter "password" was null or undefined when calling postCompanyValidateSignUp().',
      );
    }

    const queryParameters: any = {};

    const headerParameters: runtime.HTTPHeaders = {};

    const consumes: runtime.Consume[] = [{ contentType: "multipart/form-data" }];
    // @ts-ignore: canConsumeForm may be unused
    const canConsumeForm = runtime.canConsumeForm(consumes);

    let formParams: { append(param: string, value: any): any };
    let useForm = false;
    // use FormData to transmit files using content-type "multipart/form-data"
    useForm = canConsumeForm;
    if (useForm) {
      formParams = new FormData();
    } else {
      formParams = new URLSearchParams();
    }

    if (requestParameters["name"] != null) {
      formParams.append("name", requestParameters["name"] as any);
    }

    if (requestParameters["email"] != null) {
      formParams.append("email", requestParameters["email"] as any);
    }

    if (requestParameters["password"] != null) {
      formParams.append("password", requestParameters["password"] as any);
    }

    if (requestParameters["finalTaxReturn"] != null) {
      formParams.append("finalTaxReturn", requestParameters["finalTaxReturn"] as any);
    }

    const response = await this.request(
      {
        path: `/companies/validate-sign-up`,
        method: "POST",
        headers: headerParameters,
        query: queryParameters,
        body: formParams,
      },
      initOverrides,
    );

    return new runtime.JSONApiResponse(response, (jsonValue) =>
      CompanySignUpResponseFromJSON(jsonValue),
    );
  }

  /**
   * Company Validate Sign Up
   */
  async postCompanyValidateSignUp(
    requestParameters: PostCompanyValidateSignUpRequest,
    initOverrides?: RequestInit | runtime.InitOverrideFunction,
  ): Promise<CompanySignUpResponse> {
    const response = await this.postCompanyValidateSignUpRaw(requestParameters, initOverrides);
    return await response.value();
  }
}
