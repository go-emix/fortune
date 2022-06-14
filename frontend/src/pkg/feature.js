import {isLogin} from "./auth"
import ax from "./axios"
import {saveState} from "./session"

export default async function init() {
    if (!isLogin()) {
        return
    }
    const features = await ax({
        url: "system/features",
    })
    if (!features) {
        return
    }
    saveState({features: features})
}