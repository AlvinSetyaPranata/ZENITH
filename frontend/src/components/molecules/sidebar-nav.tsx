import { JSX, ParentProps } from "solid-js";
import { Icon } from "@iconify-icon/solid";

interface SidebarNavProps extends ParentProps {
  title: string;
  id: number;
  isOpenSetter: (prev: unknown) => typeof prev;
  icon?: JSX.Element;
  href?: string;
  isOpened: () => boolean;
}

export default function SidebarNav({ title, id, icon, children, isOpenSetter,isOpened }: SidebarNavProps) {


  const handleToogle = () => {
    console.log("isOpened: " + isOpened + ", id: " + id)
    if (isOpened()) {
      isOpenSetter(0)
      return
    }


    isOpenSetter(id)
  }

  return (
    <button class="rounded-md p-2 w-full hover:cursor-pointer" on:click={ handleToogle}>
      {/* upper */}
      <div class="flex justify-between items-center">
        <h3 class={`${isOpened() ? 'text-gray-300' : 'text-gray-500'} font-medium`}>{title}</h3>
        {children && (
          <Icon icon="ion:chevron-forward-sharp" class={`size-[18px] transition-transform duration-300 ease-out text-white ${isOpened() ? "rotate-90 transition-transform duration-300 ease-in text-red-500" : "text-gray-500"}`} />
        )}
      </div>
        {/* upper */}

      <div class={`mt-6 text-sm text-left px-3 overflow-y-hidden ${isOpened() ? "max-h-[500px] transition-all duration-300 ease-in" : "max-h-0 transition-all duration-100 ease-out"} space-y-4`}>
        {children}
      </div>
    </button>
  );
}
