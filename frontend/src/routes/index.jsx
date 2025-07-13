import { Button } from "@/components/ui/button";
import { Icon } from "@iconify-icon/solid";;

export default function Home() {
  return (
    <main class="text-center mx-auto text-gray-700 p-4 bg-white min-h-screen w-full">
      <Button class="bg-black hover:cursor-pointer">
        <Icon icon="material-symbols:dashboard-outline" />
        Hello
      </Button>
    </main>
  );
}
