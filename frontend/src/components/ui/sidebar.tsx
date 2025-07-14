import { ParentProps } from "solid-js";

interface SidebarProps extends ParentProps {
    isOpened: () => boolean,
    institutionName?: string,
    institutionLogoHref?: string,
}

export default function Sidebar({ isOpened, institutionName, institutionLogoHref, children} : SidebarProps) {
  return (
    <div class={`${isOpened() ? 'max-w-[20%] transition-all duration-500 ease-out' : 'max-w-0 transition-all duration-500 ease-in'} bg-black shrink-0 overflow-hidden`}>
        {/* logo section */}
        <div class="px-4 py-8 flex flex-col items-center gap-y-4 shrink-0">
            <img src={institutionLogoHref} alt="institution-logo" class="w-[80px] max-w-[80px] shrink-0" />
            <h1 class="font-semibold text-gray-200 whitespace-nowrap">{institutionName ? institutionName : "Zenith University"}</h1>
        </div>
        {/* logo section */}

        <div class="mt-8 px-4 shrink-0">
            {children}
        </div>
    </div>
  )
}
