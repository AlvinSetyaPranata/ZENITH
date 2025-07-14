import { Icon } from "@iconify-icon/solid";
import { JSX } from "solid-js";

interface ButtonProps extends JSX.ButtonHTMLAttributes<HTMLButtonElement> {
    title: string;
    isLoading:() => boolean;
    onClick?: () => void;
}

export default function Button({ title, onClick, isLoading, ...otherBtnProps } : ButtonProps) {
  return (
    <button on:click={onClick} class={`rounded-md w-[350px] py-3 text-center bg-white text-black font-medium ${isLoading() ? "hover:cursor-not-allowed" : "hover:cursor-pointer"}`} {...otherBtnProps}>
      {isLoading() ? (
        <Icon icon="mdi:loading" class="animate-spin text-2xl" />
      ) : title}
    </button>
  )
}
