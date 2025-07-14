import { createSignal, JSX } from "solid-js";

interface InputFormProps extends JSX.InputHTMLAttributes<HTMLInputElement>  {
    title: string;
}

export default function InputForm({ title, ...inpuProps} : InputFormProps) {

    const [isFocused, setIsFocused] = createSignal(false)

  return (
    <div class="flex flex-col gap-y-2 items-start">
        <label class="text-gray-300">{title}</label>
        <input on:focus={() => setIsFocused(true)} on:blur={() => setIsFocused(false)} class={`border-gray-400 w-[300px] bg-background-700 rounded-md py-2 px-3 text-sm outline-none ${isFocused() ? 'border border-white' : 'border border-transparent'}`}  {...inpuProps} />
    </div>
  )
}
