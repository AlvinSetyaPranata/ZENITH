import { Icon } from "@iconify-icon/solid";
import { createEffect } from "solid-js";
import { Motion } from "solid-motionone";


interface SoonerProps {
    title?: string;
    desc?: string;
    isClosed: () => boolean;
    setIsClosed: (state: boolean) => void;
}

export default function Sooner({ title, desc, isClosed, setIsClosed } : SoonerProps) {

  return (
     <Motion.div
      initial={{ top: "-100px", opacity: 0, left: "50%" }}
      animate={{
        top: isClosed() ? "-100px" : "20px",
        opacity: isClosed() ? 0 : 1,
      }}
      transition={{ duration: 0.5 }}
      class="fixed w-[400px] rounded-md bg-neutral-800 border border-gray-700 p-4 flex items-center justify-between"
      style={{ position: "fixed", left: "50%", transform: "translateX(-50%)" }}
    >
      <div class="space-y-0.5">
        <h2 class="text-sm">{title}</h2>
        <p class="text-sm text-gray-300 font-light">{desc}</p>
      </div>
      <Icon icon="material-symbols:close" onClick={() => setIsClosed(true)} />
    </Motion.div>
  )
}
