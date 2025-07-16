import { redirect } from "@solidjs/router"
import { useContext } from "solid-js"
import { AuthContext } from "~/contexts/auth-context"

export async function GetStudentData() {
    const authContext = useContext(AuthContext)


    const userData = authContext?.getUserData()
    
    if (!userData?.user) {
        console.error("Cannot get user?.user?.id")
        return
    }

    const studentRes = await fetch(`${import.meta.env.VITE_BASE_API_URL}/master/student?user_id=${userData?.user?.id}`, {
        headers: {
            Authorization: userData?.token ? `Bearer ${userData?.token}` : ""
        }
    })

    // Handle response

    
    if (studentRes.status == 401) {
        // throw new Error("Unauthorized")
        return false
    }

    const studentData = await studentRes.json()

    return studentData.data
}