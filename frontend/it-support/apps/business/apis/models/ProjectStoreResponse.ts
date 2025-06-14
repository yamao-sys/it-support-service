/* tslint:disable */
/* eslint-disable */
/**
 * Business Service
 * business APIs
 *
 * The version of the OpenAPI document: 1.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from "../runtime";
import type { Project } from "./Project";
import {
  ProjectFromJSON,
  ProjectFromJSONTyped,
  ProjectToJSON,
  ProjectToJSONTyped,
} from "./Project";
import type { ProjectValidationError } from "./ProjectValidationError";
import {
  ProjectValidationErrorFromJSON,
  ProjectValidationErrorFromJSONTyped,
  ProjectValidationErrorToJSON,
  ProjectValidationErrorToJSONTyped,
} from "./ProjectValidationError";

/**
 *
 * @export
 * @interface ProjectStoreResponse
 */
export interface ProjectStoreResponse {
  /**
   *
   * @type {Project}
   * @memberof ProjectStoreResponse
   */
  project: Project;
  /**
   *
   * @type {ProjectValidationError}
   * @memberof ProjectStoreResponse
   */
  errors: ProjectValidationError;
}

/**
 * Check if a given object implements the ProjectStoreResponse interface.
 */
export function instanceOfProjectStoreResponse(value: object): value is ProjectStoreResponse {
  if (!("project" in value) || value["project"] === undefined) return false;
  if (!("errors" in value) || value["errors"] === undefined) return false;
  return true;
}

export function ProjectStoreResponseFromJSON(json: any): ProjectStoreResponse {
  return ProjectStoreResponseFromJSONTyped(json, false);
}

export function ProjectStoreResponseFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): ProjectStoreResponse {
  if (json == null) {
    return json;
  }
  return {
    project: ProjectFromJSON(json["project"]),
    errors: ProjectValidationErrorFromJSON(json["errors"]),
  };
}

export function ProjectStoreResponseToJSON(json: any): ProjectStoreResponse {
  return ProjectStoreResponseToJSONTyped(json, false);
}

export function ProjectStoreResponseToJSONTyped(
  value?: ProjectStoreResponse | null,
  ignoreDiscriminator: boolean = false,
): any {
  if (value == null) {
    return value;
  }

  return {
    project: ProjectToJSON(value["project"]),
    errors: ProjectValidationErrorToJSON(value["errors"]),
  };
}
