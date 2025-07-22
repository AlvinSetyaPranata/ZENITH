/*
All internal data api should be use this base api in order to get location if provided 
*/

import { useNavigate } from "@solidjs/router"
import { useContext } from "solid-js"
import { AuthContext } from "~/contexts/auth-context"


const BASE_API_URL = import.meta.env.VITE_BASE_API_URL


export async function refreshToken() {

    const authContext = useContext(AuthContext)

    const accessTokenResponse = await fetch(`${BASE_API_URL}/token/refresh`, {
        method: 'POST',
        credentials: 'include'
    })

    if (accessTokenResponse.status != 200) {
        console.error("Error during getting new access token")
        sessionStorage.clear()
        return false
    }

    const data = await accessTokenResponse.json()

    authContext?.storeUserData(data.data, data.access_token)

    return true

}



// Return two values 
// 1. response
// 2. status flag should be logout or not

export async function baseApi(url: string, options: RequestInit) {
    const response = await fetch(`${url}`, options)

    if (response.headers.get("Location")) {
        // Token is expired
        const access_token = await refreshToken()
        if (!access_token) {
            return {status: true, response: response}
        }
    }

    return {shouldBeLogout: false, response: response}
}
