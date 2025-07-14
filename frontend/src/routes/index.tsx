import Button from "~/components/atoms/button";
import InputForm from "~/components/atoms/forms/input-form";

export default function Login() {
  return (
    <div class="bg-black w-full min-h-screen grid place-items-center">
      <div class="bg-background-800 w-1/2 h-[500px] text-center p-8">
          <h1 class="text-3xl font-semibold">Selamat Datang</h1>
          <p class="mt-4 text-gray-400">Yuk, login dulu sebelum masuk, ^-^</p>

          <div class="space-y-12 mt-18 flex flex-col items-center">
            <InputForm title="Username" />
            <InputForm title="Password" type="password" />
            <Button title="Masuk" />
          </div>
      </div>

    </div>
  )
}
