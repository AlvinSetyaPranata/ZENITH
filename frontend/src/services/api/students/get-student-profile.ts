import { useContext } from "solid-js"
import { AuthContext } from "~/contexts/auth-context"
import { baseApi } from "../base"
import { useNavigate } from "@solidjs/router"

export async function GetStudentData() {
    const authContext = useContext(AuthContext)
    const navigate = useNavigate()

    const userData = authContext?.getUserData()
    
    if (!userData?.user) {
        console.error("Cannot get user?.user?.id")
        return
    }

    const { response, shouldBeLogout} = await baseApi(`${import.meta.env.VITE_BASE_API_URL}/master/student?user_id=${userData?.user?.id}`, {
        headers: {
            Authorization: userData?.token ? `Bearer ${userData?.token}` : ""
        }
    })

    
    if (shouldBeLogout) {
        navigate("/")
        return false
    }

    const studentData = await response.json()

    return studentData.data
}