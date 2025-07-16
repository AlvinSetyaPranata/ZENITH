import { useNavigate } from "@solidjs/router";
import { createEffect, createResource, createSignal, onMount, useContext } from "solid-js";
import z, { ZodError } from "zod";
import { ZodIssue } from "zod";
import { $ZodIssue } from "zod/v4/core";
import Button from "~/components/atoms/button";
import InputForm from "~/components/atoms/forms/input-form";
import Sooner from "~/components/atoms/sooner";
import { AuthContext } from "~/contexts/auth-context";
import { SoonerContext } from "~/contexts/sooner-context";

export default function Login() {
  const navigation = useNavigate();
  const [isLoading, setIsLoading] = createSignal(false);
  const [errors, setErrors] = createSignal<$ZodIssue[]>([])
  const sooner  = useContext(SoonerContext)
  const auth = useContext(AuthContext);


  const schema = z.object({
    email: z.email({ error: "Email tidak valid"}),
    password: z.string().min(1, { error: "Harap masukkan password!"})
  })

  
  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

  
    const formData = new FormData(event.target as HTMLFormElement);
  
    const credentialData = {
      email: formData.get("email")?.toString(),
      password: formData.get("password")?.toString(),
    }


    // Perform validation
    const isFormValid = schema.safeParse(credentialData)

    const issues = isFormValid.error?.issues

    if (issues) {
      setErrors(issues.flatMap(issue => issue))
      return
    } 


    setIsLoading(true)
    
    const isLoggedIn = await auth?.login?.(credentialData);
    
    if (!isLoggedIn) {
      sooner?.show("Gagal login", "Harap Periksa kembali email dan password anda!")
      return
    }

    
    sooner?.show("Berhasil login", "Selamat datang!")
    setTimeout(() => {
      if (isLoggedIn) {
        // TODO: Later on will be overrided by role based
        setIsLoading(false)
        navigation("/student/dashboard")
        return
      }

    }, 1000)
  };


  const isInvalid = (fieldName: string, grabMessegeOnly?: boolean) => {
    const expectedField = errors().filter(fieldIssue => fieldIssue.path[0] == fieldName)


    if (grabMessegeOnly && expectedField.length >= 1) {
        return expectedField[0].message
    }

    if (expectedField.length >= 1) return true

    return false
  }


  onMount(() => {
    if (!auth) {
      throw new Error("Auth is not running!")
    }

    if (!sooner) {
      throw new Error("Sooner is not running!")
    }

  })


  return (
    <div class="bg-black w-full min-h-screen flex flex-col items-center justify-center gap-12">

      {/* institution logo */}
      <div class="w-full flex flex-col items-center">
        <img class="w-[100px]" src="https://i0.wp.com/www.uim.ac.id/uimv2/wp-content/uploads/2020/10/Ico.png?resize=200%2C200&ssl=1" alt="institute-logo" />
        <h1 class="mt-4 text-lg">Universistas Islam Madura</h1>
      </div>
      {/* institution logo */}
      <div class="bg-background-800 w-1/2 h-[500px] text-center p-8">
        <h1 class="text-3xl font-semibold">Selamat Datang</h1>
        <p class="mt-4 text-gray-400">Yuk, login dulu sebelum masuk, ^-^</p>

        <form
          on:submit={handleSubmit}
          class="space-y-12 mt-18 flex flex-col items-center"
        >
          <InputForm invalidMessege={() => isInvalid("email", true)} isInvalid={() => isInvalid("email")} title="Email" name="email" type="email" />
          <InputForm invalidMessege={() => isInvalid("password", true)} isInvalid={() => isInvalid("password")} title="Password" name="password" type="password" />
          <Button isLoading={isLoading} title="Masuk" type="submit" />
        </form>
      </div>
    </div>
  );
}
