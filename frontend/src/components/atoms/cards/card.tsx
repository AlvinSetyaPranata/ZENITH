import { Icon } from "@iconify-icon/solid";

interface CardProps {
  icon: string;
  title: string;
  desc: string;
}

export default function Card({ icon, title, desc }: CardProps) {
  return (
    <div class="max-w-[300px] h-[149px] rounded-md w-full bg-gradient-to-br to-[#1C1B1B] from-[#323232] px-4 py-5 flex flex-col justify-between">
      {/* head section */}

      <h4 class="text-gray-300 capitalize">{title}</h4>

      {/* head section */}
      <div class="w-full flex justify-between items-center">
        <Icon icon={icon} class="text-3xl text-white" />
        <h1 class="text-2xl font-semibold text-white">{desc}</h1>
      </div>
    </div>
  );
}
