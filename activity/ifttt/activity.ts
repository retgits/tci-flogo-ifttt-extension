/**
 * Imports
 */
import { Observable } from "rxjs/Observable";
import { Injectable, Injector, Inject } from "@angular/core";
import { Http } from "@angular/http";
import {
    WiContrib,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IFieldDefinition,
    IActivityContribution,
    IConnectorContribution,
    WiContributionUtils
} from "wi-studio/app/contrib/wi-contrib";

/**
 * Main
 */
@WiContrib({})
@Injectable()
export class IFTTTActivityContribution extends WiServiceHandlerContribution {
    constructor( @Inject(Injector) injector, private http: Http) {
        super(injector, http);
    }

    /**
     * The value object allows you to specify what types of values you can pick for a certain field
     */
    value = (fieldName: string, context: IActivityContribution): Observable<any> | any => {
        /**
         * For the field iftttConnection the only allowed types are connections that
         * are created as an iftttConnector (the connector category as specified in the 
         * connector.json must match what we specify here)
         */
        if (fieldName === "iftttConnection") {
            return Observable.create(observer => {
                let connectionRefs = [];
                /**
                 * The category is IFTTT
                 */
                WiContributionUtils.getConnections(this.http, "IFTTT").subscribe((data: IConnectorContribution[]) => {
                    data.forEach(connection => {
                        /**
                         * Create a list with all IFTTT connectors that have been created by the user 
                         */
                        for (let i = 0; i < connection.settings.length; i++) {
                            if (connection.settings[i].name === "name") {
                                connectionRefs.push({
                                    "unique_id": WiContributionUtils.getUniqueId(connection),
                                    "name": connection.settings[i].value
                                });
                                break;
                            }
                        }
                    });
                    observer.next(connectionRefs);
                });
            });
        }
        return null;
    }

    /**
     * The validate object can be used to validate the input of certain fields
     */
    validate = (fieldName: string, context: IActivityContribution): Observable<IValidationResult> | IValidationResult => {
        /**
         * For the field iftttConnection check that the connection has been set, otherwise
         * display the errormessage
         */
        if (fieldName === "iftttConnection") {
            let connection: IFieldDefinition = context.getField("iftttConnection")
            if (connection.value === null) {
                return ValidationResult.newValidationResult().setError("IFTTT-1000", "IFTTT Connection must be configured");
            }
        }
        return null;
    }
}