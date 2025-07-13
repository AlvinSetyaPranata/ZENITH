import { children, ParentProps } from "solid-js"

interface SidebarComponentProps extends ParentProps {
    title: string,
    logoHref: string
}



export default function Sidebar({ title, logoHref, children } : SidebarComponentProps) {
    return (
        <div class="min-h-screen min-w-max w-[25%] bg-stone-950 text-white py-6 px-4">
            {/* logo section */}
            <div class="flex flex-col items-center gap-y-4">
                <img src={logoHref} alt="logo" class="w-[100px]" />
                <p class="font-semibold text-lg text-center">{title}</p>
            </div>
            {/* logo section */}

            {/* navigation section */}
            <div class="mt-16">
                {children}
            </div>
            {/* navigation section */}
        </div>
    )
}