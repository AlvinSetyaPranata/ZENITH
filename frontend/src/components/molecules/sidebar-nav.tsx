import {  ParentProps } from "solid-js";
import { Icon } from "@iconify-icon/solid";

interface SidebarNavProps extends ParentProps {
  title: string;
  href?: string;
}

export default function SidebarNav({ title, children, href }: SidebarNavProps) {

    // const location = useLocation()

  return (
    <div class="flex justify-between items-center">
      <h3 class="text-gray-500 font-medium">{title}</h3>
      {children && (
        <Icon icon="ion:chevron-down-sharp" class="text-gray-500 text-xl" />
      )}
    </div>
  );
}
