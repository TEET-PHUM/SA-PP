/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Eatinghistory
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntMealplanEdges,
    EntMealplanEdgesFromJSON,
    EntMealplanEdgesFromJSONTyped,
    EntMealplanEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMealplan
 */
export interface EntMealplan {
    /**
     * Date holds the value of the "date" field.
     * @type {string}
     * @memberof EntMealplan
     */
    date?: string;
    /**
     * 
     * @type {EntMealplanEdges}
     * @memberof EntMealplan
     */
    edges?: EntMealplanEdges;
    /**
     * FoodID holds the value of the "food_id" field.
     * @type {number}
     * @memberof EntMealplan
     */
    foodId?: number;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntMealplan
     */
    id?: number;
    /**
     * MealID holds the value of the "meal_id" field.
     * @type {number}
     * @memberof EntMealplan
     */
    mealId?: number;
}

export function EntMealplanFromJSON(json: any): EntMealplan {
    return EntMealplanFromJSONTyped(json, false);
}

export function EntMealplanFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMealplan {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'date': !exists(json, 'date') ? undefined : json['date'],
        'edges': !exists(json, 'edges') ? undefined : EntMealplanEdgesFromJSON(json['edges']),
        'foodId': !exists(json, 'food_id') ? undefined : json['food_id'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'mealId': !exists(json, 'meal_id') ? undefined : json['meal_id'],
    };
}

export function EntMealplanToJSON(value?: EntMealplan | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'date': value.date,
        'edges': EntMealplanEdgesToJSON(value.edges),
        'food_id': value.foodId,
        'id': value.id,
        'meal_id': value.mealId,
    };
}

