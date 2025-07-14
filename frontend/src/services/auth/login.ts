import { LoginWithNimCredentialsType } from "~/types/auth-types/credentials"



export async function LoginWithNimService(credentials: LoginWithNimCredentialsType) {

    const requestBody = JSON.stringify(credentials)
    return await fetch(`${import.meta.env.VITE_BASE_PUBLIC_URL}/login`, {
        method: 'POST',
        headers: {
          'Content-Type' : 'application/json'  
        },
        body: requestBody
    })
}