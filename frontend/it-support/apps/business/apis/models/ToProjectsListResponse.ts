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
import type { ToProject } from "./ToProject";
import {
  ToProjectFromJSON,
  ToProjectFromJSONTyped,
  ToProjectToJSON,
  ToProjectToJSONTyped,
} from "./ToProject";

/**
 *
 * @export
 * @interface ToProjectsListResponse
 */
export interface ToProjectsListResponse {
  /**
   *
   * @type {Array<ToProject>}
   * @memberof ToProjectsListResponse
   */
  projects: Array<ToProject>;
  /**
   *
   * @type {string}
   * @memberof ToProjectsListResponse
   */
  nextPageToken: string;
}

/**
 * Check if a given object implements the ToProjectsListResponse interface.
 */
export function instanceOfToProjectsListResponse(value: object): value is ToProjectsListResponse {
  if (!("projects" in value) || value["projects"] === undefined) return false;
  if (!("nextPageToken" in value) || value["nextPageToken"] === undefined) return false;
  return true;
}

export function ToProjectsListResponseFromJSON(json: any): ToProjectsListResponse {
  return ToProjectsListResponseFromJSONTyped(json, false);
}

export function ToProjectsListResponseFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): ToProjectsListResponse {
  if (json == null) {
    return json;
  }
  return {
    projects: (json["projects"] as Array<any>).map(ToProjectFromJSON),
    nextPageToken: json["nextPageToken"],
  };
}

export function ToProjectsListResponseToJSON(json: any): ToProjectsListResponse {
  return ToProjectsListResponseToJSONTyped(json, false);
}

export function ToProjectsListResponseToJSONTyped(
  value?: ToProjectsListResponse | null,
  ignoreDiscriminator: boolean = false,
): any {
  if (value == null) {
    return value;
  }

  return {
    projects: (value["projects"] as Array<any>).map(ToProjectToJSON),
    nextPageToken: value["nextPageToken"],
  };
}
