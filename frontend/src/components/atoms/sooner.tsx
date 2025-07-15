import { Icon } from "@iconify-icon/solid";
import { createEffect, createSignal, onMount } from "solid-js";
import { Motion } from "solid-motionone";


interface SoonerProps {
    title: string;
    desc: string;
    isActive?: () => boolean;
    setIsActive: (state: boolean) => void;
}

export default function Sooner({ title, desc, isActive, setIsActive } : SoonerProps) {

    const [isClosed, setIsClosed] = createSignal<boolean|undefined>(true)

    onMount(() => {
        if (isClosed()) {
            setTimeout(() => setIsClosed(true), 3000)
        }
    })
    
    createEffect(() => {
        if (isActive) {
            setIsClosed(!isActive())
            setTimeout(() => setIsClosed(true), 3000)
        }
    }, [isActive])

    createEffect(() => console.log(isClosed()), [isClosed])


    const handleActive = (state: boolean) => {
        setIsClosed(state)
        setIsActive(state)
    }

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
      <Icon icon="material-symbols:close" onClick={() => handleActive(true)} />
    </Motion.div>
  )
}
