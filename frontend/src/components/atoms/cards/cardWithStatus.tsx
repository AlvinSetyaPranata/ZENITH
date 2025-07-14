import { Icon } from "@iconify-icon/solid";
import { convertTimestampToDate } from "~/utils/converters";

interface CardWithStatusProps {
  title: string;
  time: string;
  status: string;
}

export default function CardWithStatus({ title, time, status }: CardWithStatusProps) {
  return (
    <div class="max-w-[350px] h-[149px] rounded-md w-full bg-gradient-to-br to-[#1C1B1B] from-[#323232] px-4 py-5  flex flex-col justify-between">
      {/* head section */}
      <h4 class="text-gray-200 uppercase font-light">{title}</h4>
      {/* head section */}
      <div class="w-full flex justify-between items-end">
        <p class="text-gray-300">{convertTimestampToDate(time)}</p>
        <h1 class="font-semibold text-white">{status}</h1>
      </div>
    </div>
  );
}
