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



export async function LogoutService(user_id?: number) {

  if (!user_id) return false

  const res = await fetch(`${import.meta.env.VITE_BASE_PUBLIC_URL}/logout`, {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({
      id: user_id
    })
  })

  if (res.status != 200) {
    return false
  } 
  
  return true
}