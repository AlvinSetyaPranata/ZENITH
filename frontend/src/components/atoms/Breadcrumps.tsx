import { Icon } from "@iconify-icon/solid";
import { A } from "@solidjs/router";

interface BreadcrumpsProps {
  paths: { label: string; href: string }[];
}

export default function Breadcrumps({ paths }: BreadcrumpsProps) {
  return (
    <div class="flex gap-x-4 items-center px-4 py-2 w-max text-white">
      {paths.map((path) => (
        <>
          <A class="hover:underline" href={path.href}>
            {path.label}
          </A>
          {paths.length > 1 && (
            <Icon icon="ion:chevron-forward-sharp" class="size-[16px]" />
          )}
        </>
      ))}
    </div>
  );
}
