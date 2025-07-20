import { Icon } from "@iconify-icon/solid";
import { createEffect, onMount } from "solid-js";
import { Motion } from "solid-motionone";


interface SoonerProps {
    id: number;
    title?: string;
    desc?: string;
    timeout?: number;
    setIsClosed: (id: number) => void;

}

export default function Sooner({ id, title, desc, setIsClosed, timeout } : SoonerProps) {

  onMount(() => setTimeout(() => {
    setIsClosed(id)
  }, timeout ?? 3000))
 

  return (
    <Motion.div
      initial={{ top: "-100px", opacity: 0 }}
      exit={{ top: 0, opacity: 0, left: "100px" }}
      animate={{ top: 0, opacity: 1 }}
      transition={{ duration: 0.5, easing: "ease-out" }}
      class="w-[400px] rounded-md bg-neutral-800 border border-gray-700 p-4 flex items-center justify-between static"
      id={`sooner-${id}`}
    >
      <div class="space-y-0.5">
        <h2 class="text-sm">{title}</h2>
        <p class="text-sm text-gray-300 font-light">{desc}</p>
      </div>
      <Icon class="hover:cursor-pointer" icon="material-symbols:close" onClick={() => setIsClosed(id)} />
    </Motion.div>
  )
}
