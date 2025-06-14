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

/**
 *
 * @export
 * @interface ProjectsListResponse
 */
export interface ProjectsListResponse {
  /**
   *
   * @type {Array<Project>}
   * @memberof ProjectsListResponse
   */
  projects: Array<Project>;
  /**
   *
   * @type {string}
   * @memberof ProjectsListResponse
   */
  nextPageToken: string;
}

/**
 * Check if a given object implements the ProjectsListResponse interface.
 */
export function instanceOfProjectsListResponse(value: object): value is ProjectsListResponse {
  if (!("projects" in value) || value["projects"] === undefined) return false;
  if (!("nextPageToken" in value) || value["nextPageToken"] === undefined) return false;
  return true;
}

export function ProjectsListResponseFromJSON(json: any): ProjectsListResponse {
  return ProjectsListResponseFromJSONTyped(json, false);
}

export function ProjectsListResponseFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): ProjectsListResponse {
  if (json == null) {
    return json;
  }
  return {
    projects: (json["projects"] as Array<any>).map(ProjectFromJSON),
    nextPageToken: json["nextPageToken"],
  };
}

export function ProjectsListResponseToJSON(json: any): ProjectsListResponse {
  return ProjectsListResponseToJSONTyped(json, false);
}

export function ProjectsListResponseToJSONTyped(
  value?: ProjectsListResponse | null,
  ignoreDiscriminator: boolean = false,
): any {
  if (value == null) {
    return value;
  }

  return {
    projects: (value["projects"] as Array<any>).map(ProjectToJSON),
    nextPageToken: value["nextPageToken"],
  };
}
