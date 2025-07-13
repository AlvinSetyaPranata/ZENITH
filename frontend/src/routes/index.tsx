import { MetaProvider, Title } from "@solidjs/meta";
import { createSignal } from "solid-js";
import Breadcrumps from "~/components/atoms/Breadcrumps";

export default function Home() {

  const [clicked, setClicked] = createSignal(0)

  return (
    <>
    <MetaProvider>
      <Title>Zenith - Dashboard</Title>
    </MetaProvider>


    {/* content */}
    <main class="text-center mx-auto text-gray-700 p-6 bg-background-800 min-h-screen">
      <Breadcrumps paths={[{label: "Dashboard", href: "/"}]} />
    </main>
    </>
  );
}
