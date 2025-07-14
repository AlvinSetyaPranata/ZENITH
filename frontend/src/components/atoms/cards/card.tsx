import { Icon } from "@iconify-icon/solid";

interface CardProps {
    icon: string;
    title: string;
    desc: string
}

export default function Card({ icon, title, desc} : CardProps) {
  return (
    <div class="max-w-[300px] h-[149px] rounded-md w-full bg-gradient-to-br to-[#1C1B1B] from-[#323232] p-4 flex flex-col justify-between">
        {/* head section */}
        <div class="w-full flex items-center gap-x-2">
            <Icon icon={icon} class="text-2xl text-white" />
            <h4 class="text-white uppercase font-light">{title}</h4>
        </div>
        {/* head section */}
        <div class="w-full flex justify-end">
            <h1 class="text-2xl font-semibold text-white">{desc}</h1>
        </div>
    </div>
  )
}
