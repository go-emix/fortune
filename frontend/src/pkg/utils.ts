import {RouteLocationNormalizedLoaded} from "vue-router";
import {Feature} from "./session";

export function featureShow(features: Feature[], rou: RouteLocationNormalizedLoaded): any {
    let show: any = {}
    for (let i = 0; i < features.length; i++) {
        let f = features[i];
        if (f.menu_id === rou.meta.id) {
            show[f.name] = true
        }
    }
    return show
}