import { LoginWithNimCredentialsType } from "~/types/auth-types/credentials"
import { baseApi } from "../api/base"


export async function LoginWithNimService(credentials: LoginWithNimCredentialsType) {

  const requestBody = JSON.stringify(credentials)
  const {response, shouldBeLogout} =  await baseApi(`${import.meta.env.VITE_BASE_PUBLIC_URL}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: requestBody
  })

  if (shouldBeLogout) {
    sessionStorage.clear()
    return {response: response, status: false}
  }

  return {response: response, status: true}


}



export async function LogoutService(user_id?: number) {

  if (!user_id) return false

  const res = await fetch(`${import.meta.env.VITE_BASE_PUBLIC_URL}/logout`, {
    method: "POST",
    credentials: "include",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      id: user_id.toString()
    })
  })

  if (res.status != 200) {
    return false
  }

  return true
}