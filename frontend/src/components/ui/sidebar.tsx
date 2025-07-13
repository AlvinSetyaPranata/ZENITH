export interface SidebarComponentProps {
    title: string,
    logoHref: string
}



export default function Sidebar({ title, logoHref } : SidebarComponentProps) {
    return (
        <div class="min-h-screen min-w-max w-[18%] bg-stone-950 text-white py-6 px-4">
            <img src={logoHref} alt="logo" class="w-[80px]" />
            <p class="font-bold text-xl text-center font-heading">{title}</p>
        </div>
    )
}