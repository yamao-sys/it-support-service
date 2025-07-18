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

import { mapValues } from "../runtime";
/**
 *
 * @export
 * @interface CompanySignUpValidationError
 */
export interface CompanySignUpValidationError {
  /**
   *
   * @type {Array<string>}
   * @memberof CompanySignUpValidationError
   */
  name?: Array<string>;
  /**
   *
   * @type {Array<string>}
   * @memberof CompanySignUpValidationError
   */
  email?: Array<string>;
  /**
   *
   * @type {Array<string>}
   * @memberof CompanySignUpValidationError
   */
  password?: Array<string>;
  /**
   *
   * @type {Array<string>}
   * @memberof CompanySignUpValidationError
   */
  finalTaxReturn?: Array<string>;
}

/**
 * Check if a given object implements the CompanySignUpValidationError interface.
 */
export function instanceOfCompanySignUpValidationError(
  value: object,
): value is CompanySignUpValidationError {
  return true;
}

export function CompanySignUpValidationErrorFromJSON(json: any): CompanySignUpValidationError {
  return CompanySignUpValidationErrorFromJSONTyped(json, false);
}

export function CompanySignUpValidationErrorFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): CompanySignUpValidationError {
  if (json == null) {
    return json;
  }
  return {
    name: json["name"] == null ? undefined : json["name"],
    email: json["email"] == null ? undefined : json["email"],
    password: json["password"] == null ? undefined : json["password"],
    finalTaxReturn: json["finalTaxReturn"] == null ? undefined : json["finalTaxReturn"],
  };
}

export function CompanySignUpValidationErrorToJSON(json: any): CompanySignUpValidationError {
  return CompanySignUpValidationErrorToJSONTyped(json, false);
}

export function CompanySignUpValidationErrorToJSONTyped(
  value?: CompanySignUpValidationError | null,
  ignoreDiscriminator: boolean = false,
): any {
  if (value == null) {
    return value;
  }

  return {
    name: value["name"],
    email: value["email"],
    password: value["password"],
    finalTaxReturn: value["finalTaxReturn"],
  };
}
