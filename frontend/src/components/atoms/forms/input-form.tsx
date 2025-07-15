import { Icon } from "@iconify-icon/solid";
import { createEffect, createSignal, JSX } from "solid-js";

interface InputFormProps extends JSX.InputHTMLAttributes<HTMLInputElement> {
  title: string;
  isInvalid: () => boolean | string;
  invalidMessege: () => boolean | string;
}

export default function InputForm({
  title,
  isInvalid,
  invalidMessege,
  ...inputProps
}: InputFormProps) {
  const [isFocused, setIsFocused] = createSignal<boolean>(false);
  const [isPasswordShown, setIsPasswordShown] = createSignal<boolean>(false);
  const [currentStyle, setCurrentStyle] = createSignal<string>(
    "border border-transparent"
  );

  createEffect(() => {
    if (isFocused()) {
      setCurrentStyle("border border-white");
      return
    }

    if (isInvalid()) {
      setCurrentStyle("border border-red-500");
      return
    }

    setCurrentStyle("border border-transparent");
  }, [isInvalid, isFocused]);

  return (
    <div class="flex flex-col gap-y-2 items-start">
      <label class="text-gray-300">{title}</label>
      <div
        class={`flex gap-x-2 border-gray-400 w-[350px] bg-background-700 rounded-md py-2 px-3 text-sm ${currentStyle()}`}
      >
        <input
          {...inputProps}
          type={
            isPasswordShown() == false && inputProps.type == "password"
              ? "password"
              : "text"
          }
          on:focus={() => setIsFocused(true)}
          on:blur={() => setIsFocused(false)}
          class="flex-1 outline-none bg-transparent"
        />

        {inputProps.type == "password" && (
          <button
            class="hover:cursor-pointer"
            on:click={() => setIsPasswordShown((prev) => !prev)}
          >
            <Icon
              icon={
                isPasswordShown() ? "mdi:eye-outline" : "mdi:eye-off-outline"
              }
              class="text-white text-lg"
            />
          </button>
        )}
      </div>
      {isInvalid() && (
        <span class="text-xs text-red-500 font-light">{invalidMessege()}</span>
      )}
    </div>
  );
}
