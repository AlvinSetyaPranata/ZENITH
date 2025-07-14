import { useNavigate } from "@solidjs/router";
import { createResource, createSignal, useContext } from "solid-js";
import Button from "~/components/atoms/button";
import InputForm from "~/components/atoms/forms/input-form";
import { AuthContext } from "~/contexts/auth-context";
import { AuthContextType } from "~/types/auth-types/store";

export default function Login() {
  const navigation = useNavigate();
  const [isLoading, setIsLoading] = createSignal(false);


  const auth = useContext(AuthContext);
  

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    const formData = new FormData(event.target as HTMLFormElement);

    // Perform valdiation

    const isLoggedIn = await auth?.login?.({
      email: formData.get("email")?.toString(),
      password: formData.get("password")?.toString(),
    });

    if (isLoggedIn) {
      // TODO: Later on will be overrided by role based
      navigation("/student/dashboard")
      return
    }
   
    alert("Tidak bisa login")

  };


  return (
    <div class="bg-black w-full min-h-screen grid place-items-center">
      {/* institution logo */}
      <div class="w-full flex flex-col items-center">
        <img />
        <h1>Universistas Islam Madura</h1>
      </div>
      {/* institution logo */}
      <div class="bg-background-800 w-1/2 h-[500px] text-center p-8">
        <h1 class="text-3xl font-semibold">Selamat Datang</h1>
        <p class="mt-4 text-gray-400">Yuk, login dulu sebelum masuk, ^-^</p>

        <form
          on:submit={handleSubmit}
          class="space-y-12 mt-18 flex flex-col items-center"
        >
          <InputForm title="Email" name="email" type="email" />
          <InputForm title="Password" name="password" type="password" />
          <Button isLoading={isLoading} title="Masuk" type="submit" />
        </form>
      </div>
    </div>
  );
}
